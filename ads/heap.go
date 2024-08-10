package ads

import "fmt"

type Heap[K, V any] struct {
  slc []V
  valKey func(val V) K
  keyOrd func(a, b K) bool
}

func NewHeap[K, V any](
  cap int, valKey func(val V) K, keyOrd func(a, b K) bool,
) Heap[K, V] {
  return Heap[K, V]{slc: make([]V, 0, cap), valKey: valKey, keyOrd: keyOrd}
}

func (h *Heap[K, V]) Length() int {
  return len(h.slc)
}

// O(log(n))
func (h *Heap[K, V]) Push(vals ...V) {
  heapUp := func() {
    chd := len(h.slc) - 1
    par := (chd - 1) / 2
    chdKey, parKey := h.valKey(h.slc[chd]), h.valKey(h.slc[par])
    for h.keyOrd(chdKey, parKey) {
      h.slc[chd], h.slc[par] = h.slc[par], h.slc[chd]
      chd = par
      par = (chd - 1) / 2
      chdKey, parKey = h.valKey(h.slc[chd]), h.valKey(h.slc[par])
    }
  }
  for _, val := range vals {
    h.slc = append(h.slc, val)
    heapUp()
  }
}

// O(1)
func (h *Heap[K, V]) Peek() (V, error) {
  var val V
  if len(h.slc) == 0 {
    return val, fmt.Errorf("peek from empty heap")
  }
  return h.slc[0], nil
}

// O(log(n))
func (h *Heap[K, V]) Pop() (V, error) {
  child := func(par int) int {
    lft, rgh := 2 * par + 1, 2 * par + 2
    switch {
    case rgh < len(h.slc):
      lftKey, rghKey := h.valKey(h.slc[lft]), h.valKey(h.slc[rgh])
      if h.keyOrd(lftKey, rghKey) {
        return lft
      }
      return rgh
    case lft < len(h.slc):
      return lft
    default:
      return -1
    }
  }
  heapDown := func() {
    par := 0
    chd := child(par)
    if chd != -1 {
      parKey, chdKey := h.valKey(h.slc[par]), h.valKey(h.slc[chd])
      for chd != -1 && h.keyOrd(chdKey, parKey) {
        h.slc[par], h.slc[chd] = h.slc[chd], h.slc[par]
        par = chd
        chd = child(par)
      }
    }
  }
  var val V
  if len(h.slc) == 0 {
    return val, fmt.Errorf("pop from empty heap")
  }
  val, last := h.slc[0], len(h.slc) - 1
  h.slc[0] = h.slc[last]
  h.slc = h.slc[:last]
  heapDown()
  return val, nil
}
