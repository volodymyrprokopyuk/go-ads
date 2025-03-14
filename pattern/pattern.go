package pattern

import (
	"fmt"
	"time"
)

func PatDecorator() {
  type fun func(val int) int // decorated function type
  inc := func(val int) int { // decorated function
    return val + 1
  }
  logAround := func(f fun) fun { // decorator
    return func(val int) int {
      fmt.Printf("before %d\n", val)
      res := f(val)
      fmt.Printf("after %d\n", res)
      return res
    }
  }
  logAroundInc := logAround(inc)
  res := logAroundInc(1)
  fmt.Println(res) // before 1, after 2, 2
}

func Retry[R any](
  call func() (R, error), times int, delay time.Duration,
) (R, error) {
  var res R
  var err error
  for range times {
    res, err = call()
    if err == nil {
      break
    }
    time.Sleep(delay)
    delay *= 2
  }
  return res, err
}

func PatRetry() {
  calls := 0
  task := func() (int, error) {
    time.Sleep(50 * time.Millisecond)
    calls++
    fmt.Println(calls)
    if calls < 3 {
      return 0, fmt.Errorf("task error")
    }
    return calls, nil
  }
  res, err := Retry(task, 3, 500 * time.Millisecond)
  fmt.Println(res, err)
}

func PatOptions() {
  type config struct {
    value int
  }
  // In-place update of a config with the closed over value
  type option func(cfg *config) error
  configure := func(opts ...option) (*config, error) {
    var cfg config
    for _, opt := range opts { // Application of all configured options
      err := opt(&cfg)
      if err != nil {
        return nil, err
      }
    }
    return &cfg, nil
  }
  withValue := func(val int) option { // An option closure
    return func(cfg *config) error {
      if val < 0 {
        return fmt.Errorf("invalid value %d", val)
      }
      cfg.value = val
      return nil
    }
  }
  cfg, err := configure(withValue(1), withValue(2))
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(cfg.value) // 2
}
