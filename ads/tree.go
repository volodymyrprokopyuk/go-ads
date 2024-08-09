package ads

type Node[V any] struct {
  value V
  left, right *Node[V]
}

func (n *Node[V]) Value() V {
  return n.value
}

func (n *Node[V]) SetValue(val V) {
  n.value = val
}

type BSTree[K, V any] struct {
  root *Node[V]
  valKey func(val V) K
  keyOrd func (a, b K) bool
}

func NewBSTree[K, V any](
  valKey func(val V) K, keyOrd func (a, b K) bool,
) BSTree[K, V] {
  return BSTree[K, V]{valKey: valKey, keyOrd: keyOrd}
}

// sort order
func (t *BSTree[K, V]) InOrder() func(yield func(i int, nd *Node[V]) bool) {
  i, more := 0, true
  return  func(yield func(i int, nd *Node[V]) bool) {
    var inOrder func(nd *Node[V])
    inOrder = func(nd *Node[V]) {
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

// depth-first search DFS, prefix notation
func (t *BSTree[K, V]) PreOrder() func(yield func(i int, nd *Node[V]) bool) {
  i, more := 0, true
  return func(yield func(i int, nd *Node[V]) bool) {
    var preOrder func(nd *Node[V])
    preOrder = func(nd *Node[V]) {
      if nd != nil {
        if !more {
          return
        }
        more = yield(i, nd)
        i++
        preOrder(nd.left)
        preOrder(nd.right)
      }
    }
    preOrder(t.root)
  }
}

// postfix notation
func (t *BSTree[K, V]) PostOrder() func(yield func(i int, nd *Node[V]) bool) {
  i, more := 0, true
  return func(yield func(i int, nd *Node[V]) bool) {
    var postOrder func(nd *Node[V])
    postOrder = func(nd *Node[V]) {
      if nd != nil {
        postOrder(nd.left)
        postOrder(nd.right)
        if !more {
          return
        }
        more = yield(i, nd)
        i++
      }
    }
    postOrder(t.root)
  }
}

// breadth-first search BFS
func (t *BSTree[K, V]) LevelOrder() func(yield func(i int, nd *Node[V]) bool) {
  return func(yield func(i int, nd *Node[V]) bool) {
    i := 0
    var que Queue[*Node[V]]
    que.Enq(t.root)
    for que.Length() > 0 {
      nd, _ := que.Deq()
      if nd != nil {
        if !yield(i, nd) {
          break
        }
        i++
        que.Enq(nd.left); que.Enq(nd.right)
      }
    }
  }
}

func (t *BSTree[K, V]) Set(val V) {
  key := t.valKey(val)
  // recursive function expression requires variable declaration
  var set func(nd *Node[V]) *Node[V]
  set = func(nd *Node[V]) *Node[V] {
    if nd == nil {
      return &Node[V]{value: val}
    }
    ndKey := t.valKey(nd.value)
    switch { // binary search
    case t.keyOrd(key, ndKey):
      nd.left = set(nd.left)
    case t.keyOrd(ndKey, key):
      nd.right = set(nd.right)
    default:
      nd.value = val // update the node value for an existing key
    }
    return nd
  }
  t.root = set(t.root)
}

func (t *BSTree[K, V]) Get(val V) (*Node[V], bool) {
  key := t.valKey(val)
  var get func (nd *Node[V]) (*Node[V], bool)
  get = func (nd *Node[V]) (*Node[V], bool) {
    if nd == nil {
      return nil, false // the value is not found
    }
    ndKey := t.valKey(nd.value)
    switch { // binary search
    case t.keyOrd(key, ndKey):
      return get(nd.left)
    case t.keyOrd(ndKey, key):
      return get(nd.right)
    default:
      return nd, true // the value is found
    }
  }
  return get(t.root)
}
