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

func main() {
  var lst ads.DList[int]
  lst.PushHead(1, 2, 3)
  // lst.PushTail(1, 2, 3)

  // _, _ = lst.PopTail()
  // val, _ := lst.PopTail()
  // val, _ = lst.PopTail()
  // fmt.Println(val)

  var nd *ads.LNode[int]
  for _, nd = range lst.Backward() {
    if nd.Value() == 3 {
      break
    }
  }
  // lst.Insert(10, nd)
  lst.Delete(nd)

  for _, nd := range lst.Backward() {
    fmt.Printf("%v ", nd.Value())
  }
  fmt.Println()
  for _, nd := range lst.Forward() {
    fmt.Printf("%v ", nd.Value())
  }
  fmt.Println()
}
