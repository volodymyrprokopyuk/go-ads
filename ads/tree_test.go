package ads_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func TestBSTreeTraversals(t *testing.T) {
  var tree = ads.NewBSTree[int, int](
    func(val int) int { return val },
    func(a, b int) bool { return a < b },
  )
  tree.Set([]int{6, 3, 1, 2, 9, 0, 5, 4, 7, 8, 0}...)
  exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  got := make([]int, len(exp))
  for i, nd := range tree.InOrder() {
    got[i] = nd.Value()
  }
  if !SliceEqual(got, exp) {
    t.Errorf("invalid in-order: expected %v, got %v", exp, got)
  }
  exp = []int{6, 3, 1, 0, 2, 5, 4, 9, 7, 8}
  for i, nd := range tree.PreOrder() {
    got[i] = nd.Value()
  }
  if !SliceEqual(got, exp) {
    t.Errorf("invalid pre-order: expected %v, got %v", exp, got)
  }
  exp = []int{0, 2, 1, 4, 5, 3, 8, 7, 9, 6}
  for i, nd := range tree.PostOrder() {
    got[i] = nd.Value()
  }
  if !SliceEqual(got, exp) {
    t.Errorf("invalid post-order: expected %v, got %v", exp, got)
  }
  exp = []int{6, 3, 9, 1, 5, 7, 0, 2, 4, 8}
  for i, nd := range tree.LevelOrder() {
    got[i] = nd.Value()
  }
  if !SliceEqual(got, exp) {
    t.Errorf("invalid level-order: expected %v, got %v", exp, got)
  }
  // test iterator early exit
  for _, _ = range tree.InOrder() {
    break
  }
  for _, _ = range tree.PreOrder() {
    break
  }
  for _, _ = range tree.PostOrder() {
    break
  }
  for _, _ = range tree.LevelOrder() {
    break
  }
}

func TestBSTreeGetMinMax(t *testing.T) {
  var tree = ads.NewBSTree[int, int](
    func(val int) int { return val },
    func(a, b int) bool { return a < b },
  )
  tree.Set([]int{6, 3, 1, 2, 9, 0, 5, 4, 7, 8, 0}...)
  for _, exp := range []int{3, 9} {
    got, exist := tree.Get(exp)
    if !exist || got.Value() != exp {
      t.Errorf("invalid get: expected %v, got %v", exp, got)
    }
  }
  exp := 99
  _, exist := tree.Get(exp)
  if exist {
    t.Errorf("invalid get: exist non-existing value %v", exp)
  }
  exp = 0
  got := tree.Min()
  if got == nil || got.Value() != exp {
    t.Errorf("invalid min: expected %v, got %v", exp, got)
  }
  exp = 9
  got = tree.Max()
  if got == nil || got.Value() != exp {
    t.Errorf("invalid max: expected %v, got %v", exp, got)
  }
}
