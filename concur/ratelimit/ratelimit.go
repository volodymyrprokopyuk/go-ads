package ratelimit

import (
	"fmt"
	"time"
)

func New[T, U any](
  call func(req T) (U, error), maxCalls int, period time.Duration,
) func (req T) (U, error) {
  bucket := make(chan struct{}, maxCalls)
  for range maxCalls { // Fill the bucket
    bucket <- struct{}{}
  }
  tckRefill := time.NewTicker(period)
  go func() {
    for range tckRefill.C {
      for range maxCalls {
        select {
        case bucket <- struct{}{}: // Refill the bucket
        default: // The bucket is full
        }
      }
    }
  }()
  return func(req T) (U, error) {
    select {
    case <- bucket: // Perform the call if the bucket has tokens
      return call(req)
    default: // Return an error if the bucket is empty
      var res U
      return res, fmt.Errorf("rate limited %d over %s", maxCalls, period)
    }
  }
}
