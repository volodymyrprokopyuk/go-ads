package ads_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func TestStack(t *testing.T) {
  var stk ads.Stack[int]
  stk.Push(1, 2)
  gotLength, expLength := stk.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := 2
  got, _ := stk.Peek()
  if got != exp {
    t.Errorf("Invalid peek: expected %v, got %v", exp, got)
  }
  got, _ = stk.Pop()
  if got != exp {
    t.Errorf("Invalid pop: expected %v, got %v", exp, got)
  }
  got, _ = stk.Pop()
  _, err := stk.Peek()
  if err == nil {
    t.Errorf("expected peek from empty list error, got none")
  }
  _, err = stk.Pop()
  if err == nil {
    t.Errorf("expected pop from empty list error, got none")
  }
  gotLength, expLength = stk.Length(), 0
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
}
