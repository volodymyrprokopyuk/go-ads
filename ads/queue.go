package ads

type Queue[V any] struct {
  lst DList[V]
}

func (q *Queue[V]) Length() int {
  return q.lst.Length()
}

// O(1)
func (q *Queue[V]) Enq(vals ...V) {
  q.lst.PushTail(vals...)
}

// O(1)
func (q *Queue[V]) Peek() (V, error) {
  return q.lst.PeekHead()
}

// O(1)
func (q *Queue[V]) Deq() (V, error) {
  return q.lst.PopHead()
}

type Deque[V any] struct {
  lst DList[V]
}

func (d *Deque[V]) Length() int {
  return d.lst.Length()
}

// O(1)
func (d *Deque[V]) EnqFront(vals ...V) {
  d.lst.PushHead(vals...)
}

// O(1)
func (d *Deque[V]) EnqRear(vals ...V) {
  d.lst.PushTail(vals...)
}

// O(1)
func (d *Deque[V]) PeekFront() (V, error) {
  return d.lst.PeekHead()
}

// O(1)
func (d *Deque[V]) PeekRear() (V, error) {
  return d.lst.PeekTail()
}

// O(1)
func (d *Deque[V]) DeqFront() (V, error) {
  return d.lst.PopHead()
}

// O(1)
func (d *Deque[V]) DeqRear() (V, error) {
  return d.lst.PopTail()
}
