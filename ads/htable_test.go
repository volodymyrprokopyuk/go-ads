package ads_test

import (
	"strconv"
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func TestHTable(t *testing.T) {
  htb := ads.NewHTable[int, string](
    101, func (key int) string { return strconv.Itoa(key) },
    func(a, b int) bool { return a == b },
  )
  htb.Set(66, ">66")
  htb.Set(98, ">98")
  htb.Set(98, ">>98")
  gotLength, expLength := htb.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := ">66"
  got, exist := htb.Get(66)
  if !exist || got != exp {
    t.Errorf("invalid get: expected %v, got %v", exp, got)
  }
  exp = ">>98"
  got, exist = htb.Get(98)
  if !exist || got != exp {
    t.Errorf("invalid get updated: expected %v, got %v", exp, got)
  }
  key := 99
  _, exist = htb.Get(key)
  if exist {
    t.Errorf("exists non-existing key %v", key)
  }
  gotLength = 0
  for _, _ = range htb.Entries() {
    gotLength++
    if gotLength == 2 {
      break // test early exit from the iterator
    }
  }
  if gotLength != expLength {
    t.Errorf("invalid entries: expected %v, got %v", expLength, gotLength)
  }
  exp = ">66"
  got, deleted := htb.Delete(66)
  if !deleted || got != exp {
    t.Errorf("invalid delete: expected %v, got %v", exp, got)
  }
  exp = ">>98"
  got, deleted = htb.Delete(98)
  if !deleted || got != exp {
    t.Errorf("invalid delete updated: expected %v, got %v", exp, got)
  }
  _, deleted = htb.Delete(key)
  if deleted {
    t.Errorf("deleted non-existing key %v", key)
  }
  gotLength, expLength = htb.Length(), 0
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
}

var identity func(val int) int = func(val int) int { return val }

func TestHSetSetGetDelete(t *testing.T) {
  set := ads.NewHSet[int, int](11, identity)
  set.Set(1, 2)
  gotLength, expLength := set.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  i := 0
  for _, _ = range set.Entries() {
    i++
    if i == 2 {
      break // test early exit from the iterator
    }
  }
  exp := 1
  got, _ := set.Get(1)
  if got != exp {
    t.Errorf("invalid get: expected %v, got %v", exp, got)
  }
  val := 9
  _ , exist := set.Get(val)
  if exist {
    t.Errorf("exist non-existing val %v", val)
  }
  exp = 1
  got, _ = set.Delete(1)
  if got != exp {
    t.Errorf("invalid delete: expected %v, got %v", exp, got)
  }
  val = 9
  _ , deleted := set.Delete(val)
  if deleted {
    t.Errorf("deleted non-existing val %v", val)
  }
  gotLength, expLength = set.Length(), 1
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
}

func HSetValues(set ads.HSet[int, int]) []int {
  res := make([]int, 0, set.Length())
  for _, val := range set.Entries() {
    res = append(res, val)
  }
  ads.QuickSort(res, func(a, b int) bool { return a < b })
  return res
}

func TestHSetSubsetUnionIntersectDiff(t *testing.T) {
  a := ads.NewHSet[int, int](11, identity)
  b := ads.NewHSet[int, int](11, identity)
  a.Set(1, 2, 3)
  b.Set(2, 3, 4)
  exp := []int{1, 2, 3, 4}
  got := HSetValues(a.Union(b))
  if !SliceEqual(got, exp) {
    t.Errorf("invalid union: expected %v, got %v", exp, got)
  }
  exp = []int{2, 3}
  got = HSetValues(a.Intersect(b))
  if !SliceEqual(got, exp) {
    t.Errorf("invalid intersect: expected %v, got %v", exp, got)
  }
  exp = []int{1}
  got = HSetValues(a.Diff(b))
  if !SliceEqual(got, exp) {
    t.Errorf("invalid diff: expected %v, got %v", exp, got)
  }
  if a.Subset(b) {
    t.Errorf("invalid subset: false positive")
  }
  _, _ = a.Delete(1)
  if !a.Subset(b) {
    t.Errorf("invalid subset: false negative")
  }
  _, _ = b.Delete(4)
  if !a.Equal(b) {
    t.Errorf("invalid equal")
  }
}
