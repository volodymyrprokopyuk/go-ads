package ads

type TNode[K, V any] struct {
  key K
  value V
  left, right *TNode[K, V]
}

func (n *TNode[K, V]) Key() K {
  return n.key
}

func (n *TNode[K, V]) Value() V {
  return n.value
}

type BSTree[K, V any] struct {
  root *TNode[K, V]
  valKey func(val V) K
  keyOrd func (a, b K) bool
  keyEq func(a, b K) bool
}

func NewBSTree[K, V any](
  valKey func(val V) K,
  keyOrd func (a, b K) bool,
  keyEq func(a, b K) bool,
) BSTree[K, V] {
  return BSTree[K, V]{valKey: valKey, keyOrd: keyOrd, keyEq: keyEq}
}

// sort order
func (t *BSTree[K, V]) InOrder() func(yield func(i int, nd *TNode[K, V]) bool) {
  i, more := 0, true
  return  func(yield func(i int, nd *TNode[K, V]) bool) {
    var inOrder func(nd *TNode[K, V])
    inOrder = func(nd *TNode[K, V]) {
      if nd != nil {
        inOrder(nd.left)
        if !more { // handle early exit from the iterator
          return
        }
        more = yield(i, nd)
        i++
        inOrder(nd.right)
      }
    }
    inOrder(t.root)
  }
}

func (t *BSTree[K, V]) Set(val V) {
  key := t.valKey(val)
  var set func(nd *TNode[K, V]) *TNode[K, V] // recursive function expression
  set = func(nd *TNode[K, V]) *TNode[K, V] {
    if nd == nil {
      return &TNode[K, V]{key: key, value: val}
    }
    switch {
    case t.keyOrd(key, nd.key):
      nd.left = set(nd.left)
    case t.keyOrd(nd.key, key):
      nd.right = set(nd.right)
    default:
      nd.value = val // update the node value for an existing key
    }
    return nd
  }
  t.root = set(t.root)
}
