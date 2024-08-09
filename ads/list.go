package ads

import "fmt"

type Node[V any] struct {
  value V
  next, prev *Node[V] // list
  left, right *Node[V] // tree
}

func (n *Node[V]) Value() V {
  return n.value
}

func (n *Node[V]) SetValue(val V) {
  n.value = val
}

type nodeIter[V any] func(yield func(i int, nd *Node[V]) bool)

type List[V any] struct {
  head *Node[V]
  length int
}

func (l *List[V]) Length() int {
  return l.length
}

func (l *List[V]) Backward() nodeIter[V] {
  i, nd := 0, l.head
  return func(yield func(i int, nd *Node[V]) bool) {
    for nd != nil && yield(i, nd) {
      nd = nd.next
      i++
    }
  }
}

// O(1)
func (l *List[V]) Push(vals ...V) {
  for _, val := range vals {
    nd := &Node[V]{value: val}
    nd.next = l.head
    l.head = nd
    l.length++
  }
}

// O(1)
func (l *List[V]) Peek() (V, error) {
  var val V
  if l.head == nil {
    return val, fmt.Errorf("peek from empty list")
  }
  return l.head.value, nil
}

// O(1)
func (l *List[V]) Pop() (V, error) {
  var val V
  if l.head == nil {
    return val, fmt.Errorf("pop from empty list")
  }
  val = l.head.value
  l.head = l.head.next
  l.length--
  return val, nil
}

// O(n)
func (l *List[V]) Reverse() {
  var prev, next *Node[V]
  nd := l.head
  for nd != nil {
    next = nd.next // advance
    nd.next = prev // link backward
    prev = nd
    nd = next
  }
  l.head = prev
}

type DList[V any] struct {
  head, tail *Node[V]
  length int
}

func (l *DList[V]) Length() int {
  return l.length
}

func (l *DList[V]) Backward() nodeIter[V] {
  i, nd := 0, l.head
  return func(yield func(i int, nd *Node[V]) bool) {
    for nd != nil && yield(i, nd) {
      nd = nd.next
      i++
    }
  }
}

func (l *DList[V]) Forward() nodeIter[V] {
  i, nd := 0, l.tail
  return func(yield func(i int, nd *Node[V]) bool) {
    for nd != nil && yield(i, nd) {
      nd = nd.prev
      i++
    }
  }
}

// O(1)
func (l *DList[V]) PushHead(vals ...V) {
  for _, val := range vals {
    nd := &Node[V]{value: val}
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
func (l *DList[V]) PushTail(vals ...V) {
  for _, val := range vals {
    nd := &Node[V]{value: val}
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
func (l *DList[V]) PeekHead() (V, error) {
  var val V
  if l.head == nil {
    return val, fmt.Errorf("peek head from empty dlist")
  }
  return l.head.value, nil
}

// O(1)
func (l *DList[V]) PeekTail() (V, error) {
  var val V
  if l.tail == nil {
    return val, fmt.Errorf("peek tail from empty dlist")
  }
  return l.tail.value, nil
}

// O(1)
func (l *DList[V]) PopHead() (V, error) {
  var val V
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
func (l *DList[V]) PopTail() (V, error) {
  var val V
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
func (l *DList[V]) Insert(val V, after *Node[V]) {
  nd := &Node[V]{value: val}
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
func (l *DList[V]) Delete(nd *Node[V]) {
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
