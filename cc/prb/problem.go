package prb

import "fmt"

func ChSieveOfEratosthenes() {
  seq := func(n int) <-chan int {
    out := make(chan int)
    go func() {
      defer close(out)
      for val := 2; val < n; val++ {
        out <- val
      }
    }()
    return out
  }
  sieve := func(in <-chan int, n int) <-chan int {
    out := make(chan int)
    go func() {
      defer close(out)
      for val := range in {
        if val % n != 0 {
          out <- val
        }
      }
    }()
    return out
  }
  in := seq(100)
  out := sieve(in, 2)
  fmt.Printf("%v ", 2)
  for {
    prime, open := <- out
    if !open {
      break
    }
    fmt.Printf("%v ", prime)
    out = sieve(out, prime)
  }
  fmt.Println()
}
