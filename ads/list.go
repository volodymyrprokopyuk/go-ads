package ads

import "fmt"

type LNode[T any] struct {
  value T
  next, prev *LNode[T]
}

type List[T any] struct {
  head *LNode[T]
  length int
}

func (l *List[T]) Length() int {
  return l.length
}

func (l *List[T]) Backward() func(yield func(i int, val T) bool) {
  i, nd := 0, l.head
  return func(yield func(i int, val T) bool) {
    for nd != nil && yield(i, nd.value) {
      nd = nd.next
      i++
    }
  }
}

// O(1)
func (l *List[T]) Push(vals ...T) {
  for _, val := range vals {
    nd := &LNode[T]{value: val}
    nd.next = l.head
    l.head = nd
    l.length++
  }
}

// O(1)
func (l *List[T]) Peek() (T, error) {
  var val T
  if l.head == nil {
    return val, fmt.Errorf("peek from empty list")
  }
  return l.head.value, nil
}

// O(1)
func (l *List[T]) Pop() (T, error) {
  var val T
  if l.head == nil {
    return val, fmt.Errorf("pop from empty list")
  }
  val = l.head.value
  l.head = l.head.next
  l.length--
  return val, nil
}

// O(n)
func (l *List[T]) Reverse() {
  var prev, next *LNode[T]
  nd := l.head
  for nd != nil {
    next = nd.next // advance
    nd.next = prev // link backward
    prev = nd
    nd = next
  }
  l.head = prev
}
