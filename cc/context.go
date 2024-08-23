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

func CtxGracefulTermination() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  var wg sync.WaitGroup
  task := func(ctx context.Context, src <-chan int) {
    defer wg.Done()
    for {
      select {
      case <- ctx.Done(): // graceful termination
        for val := range src {
          time.Sleep(200 * time.Millisecond)
          fmt.Printf("%v graceful\n", val)
        }
        return
      case val, open := <- src: // normal processing
        if !open {
          return
        }
        time.Sleep(200 * time.Millisecond)
        fmt.Println(val)
      }
    }
  }
  src := make(chan int)
  wg.Add(1)
  go task(ctx, src)
  for val := range 7 {
    src <- val
    if val == 3 {
      cancel()
    }
  }
  close(src)
  wg.Wait()
}
