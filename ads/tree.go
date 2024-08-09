package ads

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

// sort order, infix notation a+b
func (t *BSTree[K, V]) InOrder() nodeIter[V] {
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

// depth-first search DFS, prefix notation +ab
func (t *BSTree[K, V]) PreOrder() nodeIter[V] {
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

// postfix notation ab+
func (t *BSTree[K, V]) PostOrder() nodeIter[V] {
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
func (t *BSTree[K, V]) LevelOrder() nodeIter[V] {
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

// O(log(n))
func (t *BSTree[K, V]) Set(vals ...V) {
  var set func(nd *Node[V], val V) *Node[V]
  set = func(nd *Node[V], val V) *Node[V] {
    key := t.valKey(val)
    if nd == nil {
      return &Node[V]{value: val}
    }
    ndKey := t.valKey(nd.value)
    switch { // binary search
    case t.keyOrd(key, ndKey):
      nd.left = set(nd.left, val)
    case t.keyOrd(ndKey, key):
      nd.right = set(nd.right, val)
    default:
      nd.value = val // update the node value for an existing key
    }
    return nd
  }
  for _, val := range vals {
    t.root = set(t.root, val)
  }
}

// O(log(n))
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

// O(log(n))
func (t *BSTree[K, V]) Delete(val V) bool {
  inOrderSucc := func(nd *Node[V]) *Node[V] {
    for nd.left != nil {
      nd = nd.left
    }
    return nd
  }
  var del func(nd *Node[V], val V) (*Node[V], bool)
  del = func(nd *Node[V], val V) (*Node[V], bool) {
    key := t.valKey(val)
    if nd == nil {
      return nil, false // the value is not found
    }
    var deleted bool
    ndKey := t.valKey(nd.value)
    switch { // binary search
    case t.keyOrd(key, ndKey):
      nd.left, deleted = del(nd.left, val)
    case t.keyOrd(ndKey, key):
      nd.right, deleted = del(nd.right, val)
    default: // the value is found
      if nd.left == nil && nd.right == nil {
        return nil, true // delete a leaf node
      }
      // one=child node: return the other child
      if nd.left == nil {
        return nd.right, true
      }
      if nd.right == nil {
        return nd.left, true
      }
      // two-children node
      succ := inOrderSucc(nd.right)
      nd.value = succ.value
      // delete the in-order successor
      nd.right, deleted = del(nd.right, succ.value)
    }
    return nd, deleted
  }
  var deleted bool
  t.root, deleted = del(t.root, val)
  return deleted
}

// O(log(n))
func (t *BSTree[K, V]) Min() *Node[V] {
  nd := t.root
  if nd != nil {
    for nd.left != nil {
      nd = nd.left
    }
  }
  return nd
}

// O(log(n))
func (t *BSTree[K, V]) Max() *Node[V] {
  nd := t.root
  if nd != nil {
    for nd.right != nil {
      nd = nd.right
    }
  }
  return nd
}
