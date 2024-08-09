package ads

type Stack[V any] struct {
  lst List[V]
}

func (s *Stack[V]) Length() int {
  return s.lst.Length()
}

// O(1)
func (s *Stack[V]) Push(vals ...V) {
  s.lst.Push(vals...)
}

// O(1)
func (s *Stack[V]) Peek() (V, error) {
  return s.lst.Peek()
}

// O(1)
func (s *Stack[V]) Pop() (V, error) {
  return s.lst.Pop()
}
