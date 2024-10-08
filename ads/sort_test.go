package ads_test

import (
	"cmp"
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func SliceEqual[T comparable](a, b []T) bool {
  if len(a) != len(b) {
    return false
  }
  for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}

func lt[T cmp.Ordered](a, b T) bool {
  return a < b
}

func gt[T cmp.Ordered](a, b T) bool {
  return b < a
}

func cm[T cmp.Ordered](a, b T) int {
  if a < b {
    return -1
  } else if b < a {
    return 1
  }
  return 0
}

type tcase[T any] struct {
  name string
  slc, exp []T
}

func cases() []tcase[int] {
  return []tcase[int]{
    {"empty slice", []int{}, []int{}},
    {"singleton slice", []int{1}, []int{1}},
    {"binary slice", []int{2, 1}, []int{1, 2}},
    {"sorted slice", []int{1, 2, 3}, []int{1, 2, 3}},
    {"unsorted slice", []int{3, 1, 2}, []int{1, 2, 3}},
    {
      "duplicate elements",
      []int{3, 1, 4, 9, 2, 8, 5, 6, 7, 9, 0},
      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
    },
  }
}

func TestInPlaceSort(t *testing.T) {
  sorts := []struct{
    name string
    sort func(slc []int, ord func(a, b int) bool)
  }{
    {"bubble sort", ads.BubbleSort[int]},
    {"insert sort", ads.InsertSort[int]},
    {"shell sort", ads.ShellSort[int]},
    {"select sort", ads.SelectSort[int]},
    {"quick sort", ads.QuickSort[int]},
  }
  for _, s := range sorts {
    for _, c := range cases() {
      s.sort(c.slc, lt)
      if !SliceEqual(c.slc, c.exp) {
        t.Errorf("%v: %v: expected %v, got %v", s.name, c.name, c.exp, c.slc)
      }
    }
  }
}

func TestCopySort(t *testing.T) {
  sorts := []struct{
    name string
    sort func(slc []int, ord func(a, b int) bool) []int
  }{
    {"merge sort", ads.MergeSort[int]},
    {"BSTree sort", ads.BSTSort[int]},
    {"Heap sort", ads.HeapSort[int]},
  }
  for _, s := range sorts {
    for _, c := range cases() {
      if s.name == "BSTree sort" && c.name == "duplicate elements" {
        continue // BSTree sort does not allow duplicate elements
      }
      got := s.sort(c.slc, lt)
      if !SliceEqual(got, c.exp) {
        t.Errorf("%v: %v: expected %v, got %v", s.name, c.name, c.exp, got)
      }
    }
  }
}
