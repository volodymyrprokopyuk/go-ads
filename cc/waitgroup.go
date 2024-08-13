package cc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WGroup struct {
  counter int
  cnd *sync.Cond
}

func NewWGroup() *WGroup {
  return &WGroup{cnd: sync.NewCond(new(sync.Mutex))}
}

func (g *WGroup) Add(n int) {
  g.cnd.L.Lock()
  defer g.cnd.L.Unlock()
  g.counter += n
}

func (g *WGroup) Wait() {
  g.cnd.L.Lock()
  defer g.cnd.L.Unlock()
  for g.counter > 0 {
    g.cnd.Wait()
  }
}

func (g *WGroup) Done() {
  g.cnd.L.Lock()
  defer g.cnd.L.Unlock()
  g.counter--
  if g.counter == 0 {
    g.cnd.Broadcast()
  }
}

func WGAllDone() {
  wg := NewWGroup()
  task := func(i int) {
    defer wg.Done()
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    fmt.Printf("%v done\n", i)
  }
  for i := range 5 {
    wg.Add(1)
    go task(i)
  }
  wg.Wait()
  fmt.Println("done!") // n done ... done!
}
