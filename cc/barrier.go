package cc

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
  n, done int
  cnd sync.Cond
}

func NewBarrier(n int) *Barrier {
  return &Barrier{n: n, cnd: *sync.NewCond(new(sync.Mutex))}
}

func (b *Barrier) Wait() { // wg.Done() + wg.Wait()
  b.cnd.L.Lock()
  defer b.cnd.L.Unlock()
  b.done++
  if b.done == b.n { // last goroutine done => all done
    b.done = 0 // reuse the barrier for the next sync
    b.cnd.Broadcast() // release all waiting goroutines
    return
  }
  b.cnd.Wait() // wait until all done
}

func BarSyncRounds() {
  n := 3
  bar := NewBarrier(n)
  var wg sync.WaitGroup
  task := func(i int) {
    defer wg.Done()
    for j := range 2 {
      time.Sleep(1000 * time.Millisecond)
      // round 0: n done ...; round 1: n done ...
      fmt.Printf("round %v: %v done\n", j, i)
      bar.Wait()
    }
  }
  for i := range n {
    wg.Add(1)
    go task(i)
  }
  wg.Wait()
}
