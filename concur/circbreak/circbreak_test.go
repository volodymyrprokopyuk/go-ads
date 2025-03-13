package circbreak_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/volodymyrprokopyuk/go-ads/concur/circbreak"
)

func makeCall[R any](res R, delay int) func() (R, error) {
  return func() (R, error) {
    time.Sleep(time.Duration(delay) * time.Millisecond)
    return res, nil
  }
}

func TestCircuitBreaker(t *testing.T) {
  cbr := circbreak.New[int](circbreak.Config{
    Timeout: 100 * time.Millisecond, MaxFail: 3,
    OpenInterval: 500 * time.Millisecond, MinSucc: 2,
    ResetPeriod: 500 * time.Millisecond,
  })
  for i, delay := range []int{
    // Closed => Open => HalfOpen => Open => HalfOpen => Closed
    50, 50, 50, 110, 110, 110, 10, 11, 50, 110, 10, 11, 50, 50, 50,
    // Reset and Closed
    110, 90, 110, 90, 90, 110, 90, 110, 90, 90, 90,
  } {
    call := makeCall(i, delay)
    res, err := cbr.Execute(call)
    fmt.Println(res, err)
    if delay == 11 {
      time.Sleep(510 * time.Millisecond)
    }
  }
}
