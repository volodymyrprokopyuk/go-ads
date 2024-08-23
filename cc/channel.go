package cc

import (
	"fmt"
	"sync"
	"time"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

type Channel[T any] struct {
  capSem *Semaphore // sender
  lenSem *Semaphore // receiver
  que *ads.Queue[T]
  mtx *sync.Mutex
}

func NewChannel[T any](cap int) *Channel[T] {
  return &Channel[T]{
    capSem: NewSemaphore(cap), lenSem: NewSemaphore(0),
    que: &ads.Queue[T]{}, mtx: &sync.Mutex{},
  }
}

func (ch *Channel[T]) Send(val T) {
  ch.capSem.Acquire() // backpressure, block on full capacity
  ch.mtx.Lock()
  ch.que.Enq(val)
  ch.mtx.Unlock()
  ch.lenSem.Release() // signal data available to a receiver
}

func (ch *Channel[T]) Recv() T {
  ch.capSem.Release() // avoid deadlocks
  ch.lenSem.Acquire() // once data is available, block on empty queue
  ch.mtx.Lock()
  val, _ := ch.que.Deq()
  ch.mtx.Unlock()
  return val
}

func ChSyncAsyncPipe() {
  n := 3
  ch := NewChannel[int](0) // sync channel
  // ch := NewChannel[int](n) // async channel
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    for {
      time.Sleep(1000 * time.Millisecond)
      val := ch.Recv() // enable sending capacity
      fmt.Printf("recv: %v\n", val)
      if val == n - 1 {
        break
      }
    }
  }()
  for val := range n {
    ch.Send(val)
    fmt.Printf("send %v\n", val)
  }
  // sync channel: send/recv 0, 1, 2
  // async channel: send 0, 1, 2; recv 0, 1, 2
  wg.Wait()
}

func ChEarlyExist() {
  ch, stop := make(chan int), make(chan struct{})
  var wg sync.WaitGroup
  task := func(i int) {
    defer wg.Done()
    for {
      select {
      case <- stop:
        fmt.Printf("%v: stop\n", i)
        return
      case val, open := <- ch:
        if !open {
          fmt.Printf("%v: done\n", i)
          return
        }
        time.Sleep(800 * time.Millisecond)
        fmt.Printf("%v: %v\n", i, val)
      }
    }
  }
  for i := range 3 {
    wg.Add(1)
    go task(i)
  }
  for val := range 10 {
    ch <- val
    if val == 7 {
      close(stop) // signal early exit
      break
    }
  }
  close(ch)
  wg.Wait()
}

func ChFanOutFanIn() {
  n := 3
  // each value is sent to only one channel
  src := make(chan int) // src => fan out => n pipes => fan in => sink
  task := func() <-chan int {
    res := make(chan int)
    go func() {
      defer close(res)
      for val := range src { // process workload item
        time.Sleep(800 * time.Millisecond)
        res <- val * 10
      }
    }()
    return res
  }
  fanIn := func(pipes []<-chan int) <-chan int {
    res := make(chan int)
    var wg sync.WaitGroup
    for _, pipe := range pipes {
      wg.Add(1)
      go func() { // combine workload results from each pipe
        defer wg.Done()
        for val := range pipe {
          res <- val
        }
      }()
    }
    go func() {
      wg.Wait()
      close(res)
    }()
    return res
  }
  pipes := make([]<-chan int, n)
  for i := range n { // fan out, distribute source workload
    pipes[i] = task()
  }
  sink := fanIn(pipes) // fan in, combine workload results
  go func() {
    for val := range 10 { // generate source workload
      src <- val
    }
    close(src)
  }()
  for val := range sink { // collect workload results
    fmt.Println(val)
  }
}

func ChBroadcast() {
  var wg sync.WaitGroup
  task := func(i int, src <-chan int) {
    defer wg.Done()
    for val := range src {
      time.Sleep(800 * time.Millisecond)
      fmt.Printf("%v: %v\n", i, val)
    }
  }
  broadcast := func(src <-chan int, n int) []chan int {
    pipes := make([]chan int, n)
    for i := range n { // create broadcast pipes
      pipes[i] = make(chan int)
    }
    go func() {
      defer func() {
        for _, pipe := range pipes { // close broadcast pipes
          close(pipe)
        }
      }()
      for val := range src {
        for _, pipe := range pipes { // sequential broadcast
          pipe <- val
        }
      }
    }()
    return pipes
  }
  // each value is sent to all channels
  src := make(chan int) // src => broadcast => n pipes
  pipes := broadcast(src, 3)
  for i, pipe := range pipes { // start broadcast processing
    wg.Add(1)
    go task(i, pipe)
  }
  for val := range 4 { // generate source workload
    src <- val
  }
  close(src)
  wg.Wait()
}

func ChPipeline() {
  pipe := func(src <-chan int, fun func(val int) int) <-chan int {
    res := make(chan int)
    go func() {
      defer close(res)
      for val := range src {
        time.Sleep(800 * time.Millisecond)
        res <- fun(val)
      }
    }()
    return res
  }
  src := make(chan int) // build a pipeline of goroutines
  inc := pipe(src, func(val int) int { return val + 1 })
  mul := pipe(inc, func(val int) int {return val * 10 })
  go func() {
    defer close(src)
    for i := range 5 { // generate source workload
      src <- i
    }
  }()
  for val := range mul { // collect pipeline results
    fmt.Println(val)
  }
}

func ChErrorHandling() {
  type result struct {
    err error
    res int
  }
  task := func(slc []int) <-chan result {
    res := make(chan result)
    go func() {
      defer close(res)
      for _, val := range slc {
        if val < 0 {
          res <- result{fmt.Errorf("oh"), 0}
          continue
        }
        res <- result{nil, val}
      }
    }()
    return res
  }
  res := task([]int{1, 2, -1, 3, -2})
  for val := range res {
    if val.err != nil {
      fmt.Println(val.err)
      continue
    }
    fmt.Println(val.res) // 1, 2, oh, 3, oh
  }
}

func ChTee() {
  tee := func(src <-chan int) (<-chan int, <-chan int) {
    res1, res2 := make(chan int), make(chan int)
    go func() {
      defer close(res1)
      defer close(res2)
      for val := range src {
        // per iteration channel copies to assign nil
        r1, r2 := res1, res2
        for range 2 {
          select {
          case r1 <- val:
            r1 = nil // blocks forever, lets send val to r2
          case r2 <- val:
            r2 = nil // blocks forever, lets send val to r1
          }
        }
      }
    }()
    return res1, res2
  }
  var wg sync.WaitGroup
  task := func(i int, src <-chan int) {
    defer wg.Done()
    for val := range src {
      time.Sleep(800 * time.Millisecond)
      fmt.Printf("%v: %v\n", i, val)
    }
  }
  src := make(chan int)
  res1, res2 := tee(src)
  wg.Add(2)
  go task(0, res1)
  go task(1, res2)
  for val := range 3 {
    src <- val
  }
  close(src)
  wg.Wait()
}

func ChMerge() {
  merge := func(src1 <-chan int, src2 <-chan int) <-chan int {
    mrg := make(chan int)
    go func() {
      defer close(mrg)
      for src1 != nil || src2 != nil {
        select {
        case val, open := <- src1:
          if !open {
            src1 = nil // block forever, remove this case from select
            break
          }
          time.Sleep(800 * time.Millisecond)
          mrg <- val
        case val, open := <- src2:
          if !open {
            src2 = nil // block forever, remove this case from select
            break
          }
          time.Sleep(800 * time.Millisecond)
          mrg <- val
        }
      }
    }()
    return mrg
  }
  task := func(slc []int) <-chan int {
    res := make(chan int)
    go func() {
      defer close(res)
      for _, val := range slc {
        res <- val
      }
    }()
    return res
  }
  src1 := task([]int{1, 2, 3})
  src2 := task([]int{10, 20, 30})
  mrg := merge(src1, src2)
  for val := range mrg {
    fmt.Println(val)
  }
}

func ChHeartbeat() {
  var wg sync.WaitGroup
  task := func(src <-chan int) (<-chan int, <-chan struct{}) {
    res, beat := make(chan int), make(chan struct{})
    tick := time.NewTicker(100 * time.Millisecond)
    go func() {
      defer wg.Done()
      defer close(res)
      defer close(beat)
      defer tick.Stop()
      for {
        select {
        case <- tick.C:
          select {
          case beat <- struct{}{}: // send heartbeat
          // default: // do not block if heartbeat is not read
          }
        case val, open := <- src:
          if !open {
            return
          }
          time.Sleep(200 * time.Millisecond)
          res <- val // processing
        }
      }
    }()
    return res, beat
  }
  watch := func(src <-chan int, beat <-chan struct{}) {
    defer wg.Done()
    timeout := 300 * time.Millisecond
    timer := time.NewTimer(timeout)
    defer timer.Stop()
    for {
      timer.Reset(timeout)
      select {
      case <- timer.C:
        fmt.Println("timeout")
      case <- beat:
        fmt.Println("heartbeat")
      case val, open := <- src:
        if !open {
          return
        }
        fmt.Println(val)
      }
    }
  }
  src := make(chan int)
  wg.Add(1)
  res, beat := task(src)
  wg.Add(1)
  go watch(res, beat)
  for val := range 5 {
    src <- val
  }
  close(src)
  wg.Wait()
}

func ChAsyncRateLimiter() {
  limit := 3
  bucket := make(chan struct{}, limit) // buffered bucket
  for range limit { // initial fill the bucket with limit tokens
    bucket <- struct{}{}
  }
  var wg sync.WaitGroup
  task := func(val int) {
    defer wg.Done()
    defer func() { // refill the bucket after processing a task
      bucket <- struct{}{}
    }()
    <- bucket // take a rate limited token from the bucket
    time.Sleep(800 * time.Millisecond)
    fmt.Println(val)
  }
  for val := range 10 {
    wg.Add(1)
    go task(val)
  }
  wg.Wait()
}
