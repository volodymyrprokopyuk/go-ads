package main

import (
	"fmt"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func newTree(vals []int) ads.BSTree[int, int] {
  var tree = ads.NewBSTree[int, int](
    func(val int) int { return val },
    func(a, b int) bool { return a < b },
    func(a, b int) bool { return a == b },
  )
  for _, val := range vals {
    tree.Set(val)
  }
  return tree
}

func main() {
  tree := newTree([]int{8, 1, 3, 2, 6, 0, 5, 4, 7, 9})
  for i, nd := range tree.InOrder() {
    fmt.Println(i, nd.Key(), nd.Value())
    if i == 5 {
      break
    }
  }
}
