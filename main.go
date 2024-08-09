package main

import (
	"fmt"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func newBSTree(vals []int) ads.BSTree[int, int] {
  var tree = ads.NewBSTree[int, int](
    func(val int) int { return val },
    func(a, b int) bool { return a < b },
  )
  for _, val := range vals {
    tree.Set(val)
  }
  return tree
}

func main() {
  tree := newBSTree([]int{8, 1, 3, 2, 6, 0, 5, 4, 7, 9})
  for i, nd := range tree.LevelOrder() {
    fmt.Println(i, nd.Value())
  }
  for _, val := range []int{8, 0, 9, 5, 4, 7, 99} {
    fmt.Println(tree.Get(val))
  }
}
