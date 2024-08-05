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
  return a > b
}

func main() {
  slc := []int{3, 1, 4, 9, 2, 8, 5, 6, 7, 9, 0}
  // ads.BubbleSort(slc, lt)
  // ads.InsertSort(slc, gt)
  // ads.ShellSort(slc, lt)
  ads.SelectSort(slc, gt)
  fmt.Println(slc)
}
