package ads

import "fmt"

type LNode[T any] struct {
  value T
  next, prev *LNode[T]
}

func (n *LNode[T]) Value() T {
  return n.value
}

func (n *LNode[T]) SetValue(val T) {
  n.value = val
}

type List[T any] struct {
  head *LNode[T]
  length int
}

func (l *List[T]) Length() int {
  return l.length
}

func (l *List[T]) Backward() func(yield func(i int, nd *LNode[T]) bool) {
  i, nd := 0, l.head
  return func(yield func(i int, nd *LNode[T]) bool) {
    for nd != nil && yield(i, nd) {
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

type DList[T any] struct {
  head, tail *LNode[T]
  length int
}

func (l *DList[T]) Length() int {
  return l.length
}

func (l *DList[T]) Backward() func(yield func(i int, nd *LNode[T]) bool) {
  i, nd := 0, l.head
  return func(yield func(i int, nd *LNode[T]) bool) {
    for nd != nil && yield(i, nd) {
      nd = nd.next
      i++
    }
  }
}

func (l *DList[T]) Forward() func(yield func(i int, nd *LNode[T]) bool) {
  i, nd := 0, l.tail
  return func(yield func(i int, nd *LNode[T]) bool) {
    for nd != nil && yield(i, nd) {
      nd = nd.prev
      i++
    }
  }
}

// O(1)
func (l *DList[T]) PushHead(vals ...T) {
  for _, val := range vals {
    nd := &LNode[T]{value: val}
    l.length++
    if l.head == nil {
      l.head, l.tail = nd, nd
      continue
    }
    nd.next = l.head
    l.head.prev = nd
    l.head = nd
  }
}

// O(1)
func (l *DList[T]) PushTail(vals ...T) {
  for _, val := range vals {
    nd := &LNode[T]{value: val}
    l.length++
    if l.tail == nil {
      l.head, l.tail = nd, nd
      continue
    }
    nd.prev = l.tail
    l.tail.next = nd
    l.tail = nd
  }
}

// O(1)
func (l *DList[T]) PeekHead() (T, error) {
  var val T
  if l.head == nil {
    return val, fmt.Errorf("peek head from empty dlist")
  }
  return l.head.value, nil
}

// O(1)
func (l *DList[T]) PeekTail() (T, error) {
  var val T
  if l.tail == nil {
    return val, fmt.Errorf("peek tail from empty dlist")
  }
  return l.tail.value, nil
}

// O(1)
func (l *DList[T]) PopHead() (T, error) {
  var val T
  if l.head == nil {
    return val, fmt.Errorf("pop head from empty dlist")
  }
  val = l.head.value
  l.length--
  if l.head.next == nil {
    l.head, l.tail = nil, nil
    return val, nil
  }
  l.head = l.head.next
  l.head.prev = nil
  return val, nil
}

// O(1)
func (l *DList[T]) PopTail() (T, error) {
  var val T
  if l.tail == nil {
    return val, fmt.Errorf("pop tail from empty dlist")
  }
  val = l.tail.value
  l.length--
  if l.tail.prev == nil {
    l.head, l.tail = nil, nil
    return val, nil
  }
  l.tail = l.tail.prev
  l.tail.next = nil
  return val, nil
}

// O(1)
func (l *DList[T]) Insert(val T, after *LNode[T]) {
  nd := &LNode[T]{value: val}
  l.length++
  if after == l.tail {
    nd.prev = after
    after.next = nd
    l.tail = nd
    return
  }
  nd.next = after.next
  after.next.prev = nd
  nd.prev = after
  after.next = nd
}

// O(1)
func (l *DList[T]) Delete(nd *LNode[T]) {
  l.length--
  if nd == l.head && l.head == l.tail {
    l.head, l.tail = nil, nil
    return
  }
  if nd == l.head {
    l.head = l.head.next
    l.head.prev = nil
    return
  }
  if nd == l.tail {
    l.tail = l.tail.prev
    l.tail.next = nil
    return
  }
  nd.prev.next = nd.next
  nd.next.prev = nd.prev
}
