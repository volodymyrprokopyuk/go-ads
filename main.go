package main

import (
	"cmp"
	"fmt"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

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

type A struct {
  a int
  b map[string]int
}

func main() {
  slc := []int{3, 1, 4, 9, 2, 8, 5, 6, 7, 9, 0}
  // slc := []A{{2, map[string]int{}}, {1, map[string]int{}}}

  // ads.BubbleSort(slc, lt)
  // ads.InsertSort(slc, lt)
  // ads.ShellSort(slc, lt)
  // ads.SelectSort(slc, lt)

  ads.QuickSort(slc, gt)
  // ads.QuickSort(slc, func(a, b A) bool { return a.a < b.a; })
  fmt.Println(slc)

  // fmt.Println(ads.MergeSort(slc, lt))
  // fmt.Println(ads.BinarySearch(slc, 1, cm))
  // fmt.Println(ads.BinarySearch(slc, 0, cm))
  // fmt.Println(ads.BinarySearch(slc, 9, cm))
  // fmt.Println(ads.BinarySearch(slc, 99, cm))
}
