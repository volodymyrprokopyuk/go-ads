package ads_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func TestList(t *testing.T) {
  var lst ads.List[int]
  lst.Push(1, 2)
  gotLength, expLength := lst.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := 2
  got, _ := lst.Peek()
  if got != exp {
    t.Errorf("invalid peek: expected %v, got %v", exp, got)
  }
  got, _ = lst.Pop()
  if got != exp {
    t.Errorf("invalid pop: expected %v, got %v", exp, got)
  }
  got, _ = lst.Pop()
  _, err := lst.Peek()
  if err == nil {
    t.Errorf("expected peek from empty list error, got none")
  }
  _, err = lst.Pop()
  if err == nil {
    t.Errorf("expected pop from empty list error, got none")
  }
  gotLength, expLength = lst.Length(), 0
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  lst.Push(1, 2, 3)
  lst.Reverse()
  gotList, expList := make([]int, lst.Length()), []int{1, 2, 3}
  for i, val := range lst.Backward() {
    gotList[i] = val
  }
  if !SliceEqual(gotList, expList) {
    t.Errorf("invalid reverse: expected %v, got %v", expList, gotList)
  }
}
