package ads

// O(n) non-cryptographic hash function
func djb2(str string) int {
  hash := 5381
  for _, chr := range str {
    hash = hash << 5 + hash + int(chr)
  }
  return hash
}

type entry[K, V any] struct {
  key K
  value V
}

type keyValueIter[K, V any] func(yield func(key K, val V) bool)

type HTable[K, V any] struct {
  buckets []DList[entry[K, V]]
  length int
  keyStr func(key K) string
  keyEq func(a, b K) bool
}

func NewHTable[K, V any](
  cap int, keyStr func(key K) string, keyEq func(a, b K) bool,
) HTable[K, V] {
  return HTable[K, V]{
    buckets: make([]DList[entry[K, V]], cap), keyStr: keyStr, keyEq: keyEq,
  }
}

func (t *HTable[K, V]) Length() int {
  return t.length
}

func (t *HTable[K, V]) Entries() keyValueIter[K, V] {
  i := 0
  return func(yield func(key K, val V) bool) {
    buckets: for i < cap(t.buckets) {
      bkt := t.buckets[i]
      for _, nd := range bkt.Backward() {
        e := nd.Value()
        if !yield(e.key, e.value) {
          break buckets
        }
      }
      i++
    }
  }
}

// O(1)
func (t *HTable[K, V]) Set(key K, val V) {
  idx := djb2(t.keyStr(key)) % cap(t.buckets)
  bkt := &t.buckets[idx] // bucket address for update
  ent := entry[K, V]{key, val}
  for _, nd := range bkt.Backward() { // linear search in a bucket
    e := nd.Value()
    if t.keyEq(e.key, key) {
      nd.SetValue(ent) // update an existing key
      return
    }
  }
  bkt.PushHead(ent) // add a new key
  t.length++
}

// O(1)
func (t *HTable[K, V]) Get(key K) (V, bool) {
  idx := djb2(t.keyStr(key)) % cap(t.buckets)
  bkt := t.buckets[idx]
  for _, nd := range bkt.Backward() { // linear search in a bucket
    e := nd.Value()
    if t.keyEq(e.key, key) {
      return e.value, true // the key found
    }
  }
  var val V
  return val, false // the key not found
}

// O(1)
func (t *HTable[K, V]) Delete(key K) (V, bool) {
  idx := djb2(t.keyStr(key)) % cap(t.buckets)
  bkt := &t.buckets[idx] // bucket address for deletion
  for _, nd := range bkt.Backward() { // linear search in a bucket
    e := nd.Value()
    if t.keyEq(e.key, key) {
      bkt.Delete(nd) // delete the key
      t.length--
      return e.value, true
    }
  }
  var val V
  return val, false // the key not found
}

type HSet[K comparable, V any] struct {
  htb map[K]V
  valKey func(val V) K
}

func NewHSet[K comparable, V any](cap int, valKey func(val V) K) HSet[K, V] {
  return HSet[K, V]{htb: make(map[K]V, cap), valKey: valKey}
}

func (s *HSet[K, V]) Length() int {
  return len(s.htb)
}

func (s *HSet[K, V]) Entries() keyValueIter[K, V] {
  return func (yield func (key K, val V) bool) {
    for key, val := range s.htb {
      if !yield(key, val) {
        break
      }
    }
  }
}

// O(1)
func (s *HSet[K, V]) Set(vals ...V) {
  for _, val := range vals {
    s.htb[s.valKey(val)] = val
  }
}

// O(1)
func (s *HSet[K, V]) Get(val V) bool {
  _, exist := s.htb[s.valKey(val)]
  return exist
}

// O(1)
func (s *HSet[K, V]) Delete(val V) bool {
  _, exist := s.htb[s.valKey(val)]
  if exist {
    delete(s.htb, s.valKey(val))
  }
  return exist
}

// O(n)
func (s *HSet[K, V]) Subset(set HSet[K, V]) bool {
  for _, val := range s.Entries() {
    if !set.Get(val) {
      return false
    }
  }
  return true
}

// O(n)
func (s *HSet[K, V]) Equal(set HSet[K, V]) bool {
  return s.Length() == set.Length() && s.Subset(set)
}

// O(m + n)
func (s *HSet[K, V]) Union(set HSet[K, V]) HSet[K, V] {
  res := NewHSet(s.Length() + set.Length(), s.valKey)
  for _, val := range s.Entries() {
    res.Set(val)
  }
  for _, val := range set.Entries() {
    res.Set(val)
  }
  return res
}

// O(n)
func (s *HSet[K, V]) Intersect(set HSet[K, V]) HSet[K, V] {
  res := NewHSet(s.Length(), s.valKey)
  for _, val := range s.Entries() {
    if set.Get(val) {
      res.Set(val)
    }
  }
  return res
}

// O(n)
func (s *HSet[K, V]) Diff(set HSet[K, V]) HSet[K, V] {
  res := NewHSet(s.Length(), s.valKey)
  for _, val := range s.Entries() {
    if !set.Get(val) {
      res.Set(val)
    }
  }
  return res
}
