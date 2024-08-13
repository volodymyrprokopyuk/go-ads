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
