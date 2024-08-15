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
