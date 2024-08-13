package cc

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
  permits int
  cnd *sync.Cond
}

func NewSemaphore(permits int) *Semaphore {
  return &Semaphore{permits: permits, cnd: sync.NewCond(new(sync.Mutex))}
}

func (s *Semaphore) Acquire() {
  s.cnd.L.Lock()
  defer s.cnd.L.Unlock()
  for s.permits <= 0 {
    s.cnd.Wait()
  }
  s.permits--
}

func (s *Semaphore) Release() {
  s.cnd.L.Lock()
  defer s.cnd.L.Unlock()
  s.permits++
  if s.permits > 0 {
    s.cnd.Broadcast()
  }
}

func SemConcurrencyLimit() {
  sem := NewSemaphore(3)
  var wg sync.WaitGroup
  task := func(i int) {
    defer wg.Done()
    sem.Acquire()
    defer sem.Release()
    time.Sleep(1000 * time.Millisecond)
    fmt.Printf("%v done\n", i) // at most 3 tasks every second
  }
  for i := range 10 {
    wg.Add(1)
    go task(i)
  }
  wg.Wait()
}
