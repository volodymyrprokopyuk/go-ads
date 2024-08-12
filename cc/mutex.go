package cc

import (
	"fmt"
	"sync"
)

func MtxCounter() {
  n := 99999
  var cnt int
  var mtx sync.Mutex
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    for range n {
      mtx.Lock()
      cnt++
      mtx.Unlock()
    }
  }()
  wg.Add(1)
  go func() {
    defer wg.Done()
    for range n {
      mtx.Lock()
      cnt--
      mtx.Unlock()
    }
  }()
  wg.Wait()
  fmt.Println(cnt)
}
