package ads_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func TestQueue(t *testing.T) {
  var que ads.Queue[int]
  que.Enq(1, 2)
  gotLength, expLength := que.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := 1
  got, _ := que.Peek()
  if got != exp {
    t.Errorf("invalid peek: expected %v, got %v", exp, got)
  }
  got, _ = que.Deq()
  if got != exp {
    t.Errorf("invalid deq: expected %v, got %v", exp, got)
  }
  got, _ = que.Deq()
  _, err := que.Peek()
  if err == nil {
    t.Errorf("expected peek head from empty dlist error, got none")
  }
  _, err = que.Deq()
  if err == nil {
    t.Errorf("expected pop head from empty dlist error, got none")
  }
  gotLength, expLength = que.Length(), 0
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
}

func TestDeque(t *testing.T) {
  var que ads.Deque[int]
  que.EnqRear(2)
  que.EnqFront(1)
  gotLength, expLength := que.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := 1
  got, _ := que.PeekFront()
  if got != exp {
    t.Errorf("invalid peek front: expected %v, got %v", exp, got)
  }
  got, _ = que.DeqFront()
  if got != exp {
    t.Errorf("invalid deq front: expected %v, got %v", exp, got)
  }
  exp = 2
  got, _ = que.PeekRear()
  if got != exp {
    t.Errorf("invalid peek rear: expected %v, got %v", exp, got)
  }
  got, _ = que.DeqRear()
  if got != exp {
    t.Errorf("invalid deq rear: expected %v, got %v", exp, got)
  }
  gotLength, expLength = que.Length(), 0
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
}
