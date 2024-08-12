package cc

import (
	"fmt"
	"math/rand"
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
  readers int // multiple readers
  readMtx *sync.Mutex
  writeMtx *sync.Mutex // single writer
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

// writer-preferring readers-writer mutex
type RWWMutex struct {
  readers int // multiple readers
  waitingWriters int // prefer waiting writers readers
  writerActive bool
  cnd *sync.Cond
}

func NewRWWMutex() *RWWMutex {
  return &RWWMutex{cnd: sync.NewCond(new(sync.Mutex))}
}

func (m *RWWMutex) Lock() {
  m.cnd.L.Lock()
  defer m.cnd.L.Unlock()
  m.waitingWriters++
  // wait for all readers and an active writer
  for m.readers > 0 || m.writerActive {
    m.cnd.Wait()
  }
  m.waitingWriters--
  m.writerActive = true // single writer
}

func (m *RWWMutex) Unlock() {
  m.cnd.L.Lock()
  defer m.cnd.L.Unlock()
  m.writerActive = false
  m.cnd.Broadcast() // single writer
}

func (m *RWWMutex) RLock() {
  m.cnd.L.Lock()
  defer m.cnd.L.Unlock()
  // wait for and active writer and all waiting writers
  // prefer waiting writers over readers
  for m.waitingWriters > 0 || m.writerActive {
    m.cnd.Wait()
  }
  m.readers++
}

func (m *RWWMutex) RUnlock() {
  m.cnd.L.Lock()
  defer m.cnd.L.Unlock()
  m.readers--
  if m.readers == 0 {
    m.cnd.Broadcast() // the last reader broadcasts
  }
}

func RWWMutexPrefersWriters() {
  mtx := NewRWWMutex()
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
    time.Sleep(10 * time.Microsecond)
    fmt.Println("read")
  }
  tasks := []func(){reader, reader, writer, writer, reader, writer}
  for _, task := range tasks {
    wg.Add(1)
    go task() // always write write write read read read
  }
  wg.Wait()
}

func CndAllJoined() {
  n := 4
  var joined int
  cnd := sync.NewCond(new(sync.Mutex))
  var wg sync.WaitGroup
  join := func(i int) {
    defer wg.Done()
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    cnd.L.Lock()
    defer cnd.L.Unlock()
    joined++
    fmt.Printf("%v joined\n", i)
    if joined == n {
      cnd.Broadcast()
    }
    for joined < n {
      cnd.Wait()
    }
    fmt.Printf("%v all joined\n", i)
  }
  for i := range n {
    wg.Add(1)
    go join(i)
  }
  wg.Wait()
}
