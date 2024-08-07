package ads

type Stack[T any] struct {
  lst List[T]
}

func (s *Stack[T]) Length() int {
  return s.lst.Length()
}

// O(1)
func (s *Stack[T]) Push(vals ...T) {
  s.lst.Push(vals...)
}

// O(1)
func (s *Stack[T]) Peek() (T, error) {
  return s.lst.Peek()
}

// O(1)
func (s *Stack[T]) Pop() (T, error) {
  return s.lst.Pop()
}
