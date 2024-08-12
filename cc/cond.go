package cc

import (
	"fmt"
	"sync"
	"time"
)

func CndBalance() {
  var balance int
  cnd := sync.NewCond(new(sync.Mutex)) // a condition requires a mutex
  listen := func(amount int) {
    cnd.L.Lock()
    // critical section 1: wait until the condition is met
    for balance < amount { // check the condition
      // implicit unlock => wait => implicit lock
      // on cnd.Signal() or cnd.Broadcast() from other goroutine
      cnd.Wait() // wait in a loop inside a critical section 1
    }
    // critical section 2: process when the condition is met
    fmt.Println(balance)
    cnd.L.Unlock()
  }
  go listen(3) // 3
  go listen(5) // 5
  for range 7 {
    time.Sleep(10 * time.Millisecond)
    cnd.L.Lock()
    balance++ // update condition state in a critical section
    // signal to a single goroutine, broadcast to multiple goroutines
    cnd.Broadcast()
    cnd.L.Unlock()
  }
}
