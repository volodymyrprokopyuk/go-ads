package concur

import (
	"fmt"
	"sync"
	"time"
)

type State string

const (
  stClosed = State("Closed")
  stOpen = State("Open")
  stHalfOpen = State("HalfOpen")
)

type circuitBreaker[R any] struct {
  timeout time.Duration // The timeout of an external call
  maxFail int // The number of failures before the Open state
  openInterval time.Duration // The duration before the HalfOpen state
  minSucc int // The number of success calls before the Closed state
  resetPeriod time.Duration // The duration before the reset in the Closed state
  mtx sync.RWMutex
  state State // The state of the circuit breaker
  numFail int // The current number of failures
  numSucc int // The current number of success calls
  tmrReset *time.Timer
}

func NewCircuitBreaker[R any](
  timeout time.Duration, maxFail int, openInterval time.Duration, minSucc int,
  resetPeriod time.Duration,
) *circuitBreaker[R] {
  return &circuitBreaker[R]{
    timeout: timeout, maxFail: maxFail, openInterval: openInterval,
    minSucc: minSucc, resetPeriod: resetPeriod, state: stClosed,
  }
}

func (c *circuitBreaker[R]) stateClosed() {
  c.state = stClosed
  c.numFail, c.numSucc = 0, 0
  c.tmrReset = time.AfterFunc(c.resetPeriod, func() {

  })
}

func (c *circuitBreaker[R]) stateOpen() {
  c.state = stOpen
  c.numFail, c.numSucc = 0, 0
  var tmrOpen *time.Timer // Open => HalfOpen
  tmrOpen = time.AfterFunc(c.openInterval, func() {
    defer tmrOpen.Stop()
    c.mtx.Lock()
    defer c.mtx.Unlock()
    c.state = stHalfOpen
    c.numFail, c.numSucc = 0, 0
  })
}

func (c *circuitBreaker[R]) Execute(call func() (R, error)) (R, error) {
  var res R

  c.mtx.RLock()
  if c.state == stOpen {
    c.mtx.RUnlock()
    return res, fmt.Errorf("circuit breaker is open")
  }
  c.mtx.RUnlock()

  succ, fail := make(chan R), make(chan error)
  defer close(fail)
  defer close(succ)
  go func() {
    res, err := call()
    if err != nil {
      fail <- err
      return
    }
    succ <- res
  }()

  var err error
  select {
  case <- time.After(c.timeout):
    c.numFail++
    err = fmt.Errorf("timeout after %s", c.timeout)
  case res = <- succ:
    c.minSucc++
  case err = <- fail:
    c.numFail++
  }

  c.mtx.Lock()
  defer c.mtx.Unlock()
  if c.state == stClosed && c.numFail >= c.maxFail { // Closed => Open
    c.stateOpen()
  }
  if c.state == stHalfOpen && c.numFail > 0 { // HalfOpen => Open
    c.stateOpen()
  }
  if c.state == stHalfOpen && c.numSucc > c.minSucc { // HalfOpen => Closed
    c.stateClosed()
  }
  return res, err
}

func CbrExecute() {
  call := func() (int, error) {
    time.Sleep(500 * time.Millisecond)
    return 1, nil
  }
  cbr := NewCircuitBreaker[int](
    100 * time.Millisecond, 3, 1000 * time.Millisecond, 2,
    5000 * time.Millisecond,
  )
  res, err := cbr.Execute(call)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(res)
}
