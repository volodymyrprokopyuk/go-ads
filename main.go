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
}
