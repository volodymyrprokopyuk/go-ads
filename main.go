package main

import (
	"fmt"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

// var heapSlice = []int{6, 3, 1, 2, 9, 0, 5, 4, 7, 8, 0}

func main() {
  trie := ads.NewTrie()
  trie.Set("go", "goal")
  fmt.Println(trie.Get("go"), trie.Get("goal"), trie.Get("goals"))

  // var heap = ads.NewHeap[int, int](
  //   11, func(val int) int { return val },
  //   func(a, b int) bool { return a < b },
  // )
  // heap.Push(heapSlice...)

  // fmt.Println(heap.Peek())
  // fmt.Println(heap.Length())

  // for heap.Length() > 0 {
  //   fmt.Println(heap.Pop())
  // }

  // fmt.Println(tree.Delete(6))

  // for i, nd := range tree.InOrder() {
  //   fmt.Println(i, nd.Value())
  // }

  // for _, val := range []int{8, 0, 9, 5, 4, 7, 99} {
  //   fmt.Println(tree.Get(val))
  // }

  // fmt.Println(tree.Min())
  // fmt.Println(tree.Max())
}
