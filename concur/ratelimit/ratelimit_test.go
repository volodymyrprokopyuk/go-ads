package ratelimit_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/volodymyrprokopyuk/go-ads/concur/ratelimit"
)

func call(i int) (int, error) {
  time.Sleep(50 * time.Millisecond)
  return i, nil
}

func TestRateLimiter(t *testing.T) {
  rlCall := ratelimit.New(call, 5, 900 * time.Millisecond)
  for i := range 10 {
    res, err := rlCall(i)
    fmt.Println(res, err)
    time.Sleep(90 * time.Millisecond)
  }
}
