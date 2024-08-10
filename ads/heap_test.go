package ads_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func newHeap() ads.Heap[int,int]{
  return ads.NewHeap[int, int](
    11, func(val int) int { return val },
    func(a, b int) bool { return a < b },
  )
}

var heapSlice = []int{6, 3, 1, 2, 9, 0, 5, 4, 7, 8, 0}

func TestHeap(t *testing.T) {
  heap := newHeap()
  heap.Push(heapSlice...)
  gotLength, expLength := heap.Length(), len(heapSlice)
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := 0
  got, _ := heap.Peek()
  if got != exp {
    t.Errorf("invalid peek: expected %v, got %v", exp, got)
  }
  expSlc := []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  gotSlc := make([]int, 0, heap.Length())
  for heap.Length() > 0 {
    val, _ := heap.Pop()
    gotSlc = append(gotSlc, val)
  }
  if !SliceEqual(gotSlc, expSlc) {
    t.Errorf("invalid pop: expected %v, got %v", expSlc, gotSlc)
  }
  _, err := heap.Peek()
  if err == nil {
    t.Errorf("invalid peek: expected peek from empty heap error, got none")
  }
  _, err = heap.Pop()
  if err == nil {
    t.Errorf("invalid pop: expected pop from empty heap error, got none")
  }
  gotLength, expLength = heap.Length(), 0
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
}
