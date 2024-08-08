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

type HTable[K, V any] struct {
  buckets []DList[entry[K, V]]
  length int
  keyStr func(key K) string
  keyEq func(a, b K) bool
}

func NewHTable[K, V any](
  cap int, keyStr func(key K) string, keyEq func(a, b K) bool,
) *HTable[K, V] {
  return &HTable[K, V]{
    buckets: make([]DList[entry[K, V]], cap), keyStr: keyStr, keyEq: keyEq,
  }
}

func (t *HTable[K, V]) Length() int {
  return t.length
}

func (t *HTable[K, V]) Entries() func(yield func(key K, val V) bool) {
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
