package concur

import (
	"fmt"
	"sync"
	"time"
)

type SemaphoreCh chan struct{}

func NewSemaphoreCh(permits int) *SemaphoreCh {
  sem := SemaphoreCh(make(chan struct{}, permits))
  for range permits {
    sem <- struct{}{}
  }
  return &sem
}

func (s *SemaphoreCh) Acquire(t time.Duration) error {
  select {
  case <- *s:
    return nil
  case <- time.After(t):
    return fmt.Errorf("timeout")
  }
}

func (s *SemaphoreCh) Release() {
  *s <- struct{}{}
}

func SemChConcurLimit() {
  sem := NewSemaphoreCh(3)
  defer close(*sem)
  var wg sync.WaitGroup
  task := func(i int) {
    defer wg.Done()
    err := sem.Acquire(1000 * time.Millisecond)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer sem.Release()
    time.Sleep(400 * time.Millisecond)
    fmt.Printf("%d done\n", i)
  }
  for i := range 10 {
    wg.Add(1)
    go task(i)
  }
  wg.Wait()
}

type SemaphoreCnd struct {
  permits int
  cnd *sync.Cond
}

func NewSemaphoreCnd(permits int) *SemaphoreCnd {
  return &SemaphoreCnd{permits: permits, cnd: sync.NewCond(new(sync.Mutex))}
}

func (s *SemaphoreCnd) Acquire() {
  s.cnd.L.Lock()
  defer s.cnd.L.Unlock()
  for s.permits <= 0 {
    s.cnd.Wait()
  }
  s.permits--
}

func (s *SemaphoreCnd) Release() {
  s.cnd.L.Lock()
  defer s.cnd.L.Unlock()
  s.permits++
  if s.permits > 0 {
    s.cnd.Broadcast()
  }
}

func SemCndConcurLimit() {
  sem := NewSemaphoreCnd(3)
  var wg sync.WaitGroup
  task := func(i int) {
    defer wg.Done()
    sem.Acquire()
    defer sem.Release()
    time.Sleep(1000 * time.Millisecond)
    fmt.Printf("%d done\n", i) // at most 3 tasks every second
  }
  for i := range 10 {
    wg.Add(1)
    go task(i)
  }
  wg.Wait()
}
