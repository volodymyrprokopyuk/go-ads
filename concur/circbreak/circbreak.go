package circbreak

import (
	"fmt"
	"sync"
	"time"
)

type state string

const (
  stClosed = state("Closed")
  stOpen = state("Open")
  stHalfOpen = state("HalfOpen")
)

type Config struct {
  Timeout time.Duration // The timeout for the external call
  MaxFail int // The number of failures before Closed => Open
  OpenInterval time.Duration // The duration before Open => HalfOpen
  MinSucc int // The number of successful calls before HalfOpen => Closed
  ResetPeriod time.Duration // The duration before the reset in the Closed state
}

type CircuitBreaker[R any] struct {
  cfg Config
  mtx sync.RWMutex // Sync access to the state from concurrent Execute calls
  state state // The state of the circuit breaker
  cntFail int // The count of failures since the last reset
  cntSucc int // The count of successful calls since the last reset
  tckReset *time.Ticker // Periodic reset of the failure/success counts
  tmrOpen *time.Timer // The trigger to move from Open => HalfOpen
}

func New[R any](cfg Config) *CircuitBreaker[R] {
  c := &CircuitBreaker[R]{cfg: cfg}
  c.state = stClosed // The initial state is Closed
  c.tckReset = time.NewTicker(c.cfg.ResetPeriod)
  go c.cntReset()
  return c
}

func (c *CircuitBreaker[R]) cntReset() {
  for range c.tckReset.C {
    c.mtx.Lock()
    if c.state == stClosed {
      fmt.Println("=> Reset")
      c.cntFail, c.cntSucc = 0, 0
    }
    c.mtx.Unlock()
  }
}

func (c *CircuitBreaker[R]) stateClosed() {
  fmt.Println("=> Closed")
  c.state = stClosed
  c.cntFail, c.cntSucc = 0, 0
  c.tckReset.Reset(c.cfg.ResetPeriod)
}

func (c *CircuitBreaker[R]) stateOpen() {
  fmt.Println("=> Open")
  c.state = stOpen
  c.cntFail, c.cntSucc = 0, 0
  c.tmrOpen = time.AfterFunc(c.cfg.OpenInterval, c.stateHalfOpen)
}

func (c *CircuitBreaker[R]) stateHalfOpen() {
  fmt.Println("=> HalfOpen")
  c.tmrOpen.Stop()
  c.mtx.Lock()
  defer c.mtx.Unlock()
  c.state = stHalfOpen
  c.cntFail, c.cntSucc = 0, 0
}

func (c *CircuitBreaker[R]) Execute(call func() (R, error)) (R, error) {
  var res R
  // Immediately return an error when in the Open state
  c.mtx.RLock()
  if c.state == stOpen {
    c.mtx.RUnlock()
    return res, fmt.Errorf("circuit breaker is open")
  }
  c.mtx.RUnlock()
  // Execute the external call in a dedicated goroutine
  succ, fail := make(chan R), make(chan error)
  go func() {
    defer close(succ)
    defer close(fail)
    res, err := call()
    if err != nil {
      fail <- err
      return
    }
    succ <- res
  }()
  // Wait for the external call success, a failure, or a timeout
  var err error
  var cntFail, cntSucc int
  select {
  case <- time.After(c.cfg.Timeout):
    cntFail++
    err = fmt.Errorf("timeout after %s", c.cfg.Timeout)
  case res = <- succ:
    cntSucc++
  case err = <- fail:
    cntFail++
  }
  // Transition to the right state
  c.mtx.Lock()
  defer c.mtx.Unlock()
  c.cntFail += cntFail
  c.cntSucc += cntSucc
  if c.state == stClosed && c.cntFail >= c.cfg.MaxFail { // Closed => Open
    c.stateOpen()
  }
  if c.state == stHalfOpen && c.cntFail > 0 { // HalfOpen => Open
    c.stateOpen()
  }
  if c.state == stHalfOpen && c.cntSucc >= c.cfg.MinSucc { // HalfOpen => Closed
    c.stateClosed()
  }
  return res, err
}
