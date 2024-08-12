package cc

import (
	"fmt"
	"sync"
	"time"
)

func MtxCounter() {
  n := 99999
  var cnt int // shared memory
  var mtx sync.Mutex
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    for range n {
      mtx.Lock()
      cnt++ // critical section
      mtx.Unlock()
    }
  }()
  wg.Add(1)
  go func() {
    defer wg.Done()
    for range n {
      mtx.Lock()
      cnt-- // critical section
      mtx.Unlock()
    }
  }()
  wg.Wait()
  fmt.Println(cnt) // 0
}

// read-preferring readers-writer mutex
type RRWMutex struct {
  readers int
  readMtx *sync.Mutex
  writeMtx *sync.Mutex
}

func NewRRWMutex() *RRWMutex {
  return &RRWMutex{readMtx: new(sync.Mutex), writeMtx: new(sync.Mutex)}
}

func (m *RRWMutex) Lock() {
  m.writeMtx.Lock() // only one exclusive writer
}

func (m *RRWMutex) Unlock() {
  m.writeMtx.Unlock()
}

func (m *RRWMutex) RLock() {
  m.readMtx.Lock()
  m.readers++
  if m.readers == 1 {
    m.writeMtx.Lock() // lock a writer while multiple readers
  }
  m.readMtx.Unlock() // allow multiple readers, prefer readers over a writer
}

func (m *RRWMutex) RUnlock() {
  m.readMtx.Lock()
  m.readers--
  if m.readers == 0 {
    m.writeMtx.Unlock() // allow a writer when no readers
  }
  m.readMtx.Unlock() // allow multiple readers, prefer readers over a writer
}

func RRWMutexPrefersReaders() {
  mtx := NewRRWMutex()
  var wg sync.WaitGroup
  writer := func() {
    defer wg.Done()
    mtx.Lock()
    defer mtx.Unlock()
    time.Sleep(10 * time.Millisecond)
    fmt.Println("write")
  }
  reader := func() {
    defer wg.Done()
    mtx.RLock()
    defer mtx.RUnlock()
    time.Sleep(10 * time.Millisecond)
    fmt.Println("read")
  }
  tasks := []func(){writer, writer, reader, reader, writer, reader}
  for _, task := range tasks {
    wg.Add(1)
    go task() // always read, read, read, write, write, write
  }
  wg.Wait()
}
