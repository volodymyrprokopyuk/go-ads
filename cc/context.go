package cc

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func CtxCancelTimeout() {
  var wg sync.WaitGroup
  task := func(ctx context.Context, i int) {
    defer wg.Done()
    for {
      select {
      case <- ctx.Done():
        switch ctx.Err() {
        case context.Canceled:
          fmt.Println("canceled")
        case context.DeadlineExceeded:
          fmt.Println("timeout")
        }
        return
      default: // non-blocking, immediate execution
        time.Sleep(200 * time.Millisecond)
        fmt.Println("working...")
      }
    }
  }
  // * cancel context
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel() // a context must be canceled
  for i := range 3 {
    wg.Add(1)
    go task(ctx, i)
  }
  time.Sleep(300 * time.Microsecond)
  cancel() // explicit cancellation
  wg.Wait()
  // * timeout context
  ctx, cancelTimeout := context.WithTimeout(
    context.Background(), 300 * time.Millisecond,
  )
  defer cancelTimeout()
  for i := range 3 {
    wg.Add(1)
    go task(ctx, i)
  }
  wg.Wait()
}
