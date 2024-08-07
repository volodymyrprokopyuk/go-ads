package ads

type Queue[T any] struct {
  lst DList[T]
}

func (q *Queue[T]) Length() int {
  return q.lst.Length()
}

// O(1)
func (q *Queue[T]) Enq(vals ...T) {
  q.lst.PushTail(vals...)
}

// O(1)
func (q *Queue[T]) Peek() (T, error) {
  return q.lst.PeekHead()
}

// O(1)
func (q *Queue[T]) Deq() (T, error) {
  return q.lst.PopHead()
}

type Deque[T any] struct {
  lst DList[T]
}

func (d *Deque[T]) Length() int {
  return d.lst.Length()
}

// O(1)
func (d *Deque[T]) EnqFront(vals ...T) {
  d.lst.PushHead(vals...)
}

// O(1)
func (d *Deque[T]) EnqRear(vals ...T) {
  d.lst.PushTail(vals...)
}

// O(1)
func (d *Deque[T]) PeekFront() (T, error) {
  return d.lst.PeekHead()
}

// O(1)
func (d *Deque[T]) PeekRear() (T, error) {
  return d.lst.PeekTail()
}

// O(1)
func (d *Deque[T]) DeqFront() (T, error) {
  return d.lst.PopHead()
}

// O(1)
func (d *Deque[T]) DeqRear() (T, error) {
  return d.lst.PopTail()
}
