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
  for i, nd := range lst.Backward() {
    gotList[i] = nd.Value()
  }
  if !SliceEqual(gotList, expList) {
    t.Errorf("invalid reverse: expected %v, got %v", expList, gotList)
  }
}

func TestDListPushPeekPopHead(t *testing.T) {
  var lst ads.DList[int]
  lst.PushHead(1, 2)
  gotLength, expLength := lst.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := 2
  got, _ := lst.PeekHead()
  if got != exp {
    t.Errorf("invalid peek head: expected %v, got %v", exp, got)
  }
  got, _ = lst.PopHead()
  if got != exp {
    t.Errorf("invalid pop head: expected %v, got %v", exp, got)
  }
  got, _ = lst.PopHead()
  _, err := lst.PeekHead()
  if err == nil {
    t.Errorf("expected peek head from empty list error, got none")
  }
  _, err = lst.PopHead()
  if err == nil {
    t.Errorf("expected pop head from empty list error, got none")
  }
}

func TestDListPushPeekPopTail(t *testing.T) {
  var lst ads.DList[int]
  lst.PushTail(1, 2)
  gotLength, expLength := lst.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := 2
  got, _ := lst.PeekTail()
  if got != exp {
    t.Errorf("invalid peek tail: expected %v, got %v", exp, got)
  }
  got, _ = lst.PopTail()
  if got != exp {
    t.Errorf("invalid pop tail: expected %v, got %v", exp, got)
  }
  got, _ = lst.PopTail()
  _, err := lst.PeekTail()
  if err == nil {
    t.Errorf("expected peek tail from empty list error, got none")
  }
  _, err = lst.PopTail()
  if err == nil {
    t.Errorf("expected pop tail from empty list error, got none")
  }
}

func DListBackward[T any](lst *ads.DList[T]) []T {
  slc := make([]T, lst.Length())
  for i, nd := range lst.Backward() {
    slc[i] = nd.Value()
  }
  return slc
}

func DListForward[T any](lst *ads.DList[T]) []T {
  slc := make([]T, lst.Length())
  for i, nd := range lst.Forward() {
    slc[i] = nd.Value()
  }
  return slc
}

func TestDListBackwardForward(t *testing.T) {
  var lst ads.DList[int]
  lst.PushHead(3, 4)
  lst.PushTail(2, 1)
  exp := []int{4, 3, 2, 1}
  got := DListBackward(&lst)
  if !SliceEqual(got, exp) {
    t.Errorf("invalid backward: expected %v, got %v", exp, got)
  }
  exp = []int{1, 2, 3, 4}
  got = DListForward(&lst)
  if !SliceEqual(got, exp) {
    t.Errorf("invalid forward: expected %v, got %v", exp, got)
  }
}

func TestDListInsert(t *testing.T) {
  cases := []struct{
    name string
    vals []int
    ndVal int
    insVal int
    exp []int
  }{
    {"insert after middle node", []int{1, 2}, 2, 10, []int{2, 10, 1}},
    {"insert after tail", []int{1, 2}, 1, 10, []int{2, 1, 10}},
  }
  for _, c := range cases {
    var lst ads.DList[int]
    lst.PushHead(c.vals...)
    var nd *ads.LNode[int]
    for _, nd = range lst.Backward() {
      if nd.Value() == c.ndVal {
        break
      }
    }
    lst.Insert(c.insVal, nd)
    got := DListBackward(&lst)
    if !SliceEqual(got, c.exp) {
      t.Errorf("invalid insert: %v: expected %v, got %v", c.name, c.exp, got)
    }
  }
}

func TestDListDelete(t *testing.T) {
  cases := []struct{
    name string
    vals []int
    ndVal int
    exp []int
  }{
    {"delete last node", []int{1}, 1, []int{}},
    {"delete head", []int{1, 2, 3}, 3, []int{2, 1}},
    {"delete tail", []int{1, 2, 3}, 1, []int{3, 2}},
    {"delete middle node", []int{1, 2, 3}, 2, []int{3, 1}},
  }
  for _, c := range cases {
    var lst ads.DList[int]
    lst.PushHead(c.vals...)
    var nd *ads.LNode[int]
    for _, nd = range lst.Backward() {
      if nd.Value() == c.ndVal {
        break
      }
    }
    lst.Delete(nd)
    got := DListBackward(&lst)
    if !SliceEqual(got, c.exp) {
      t.Errorf("invalid delete: %v: expected %v, got %v", c.name, c.exp, got)
    }
  }
}
