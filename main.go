package main

import (
	"cmp"
	"fmt"
	"strconv"

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

func djb2(str string) int {
  hash := 5381
  for _, chr := range str {
    hash = hash << 5 + hash + int(chr)
  }
  return hash
}

func collision() {
  n := 99
  seen := make(map[int]string, n)
  for i := range n {
    key := strconv.Itoa(i)
    hash := djb2(key) % 101
    if k, exist := seen[hash]; exist {
      fmt.Printf("%v, %v => %v\n", k, key, hash)
    }
    seen[hash] = key
  }
}

func main() {
  // collision()

  htb := ads.NewHTable[int, string](
    101, func (a int) string { return strconv.Itoa(a) },
    func(a, b int) bool { return a == b },
  )

  htb.Set(66, ">66")
  htb.Set(98, ">98")
  htb.Set(98, ">>98")

  fmt.Println(htb.Length())
  for k, v := range htb.Entries() {
    fmt.Println(k, v)
  }

  fmt.Println(htb.Get(66))
  fmt.Println(htb.Get(98))
  fmt.Println(htb.Get(99))

  fmt.Println(htb.Delete(66))
  fmt.Println(htb.Delete(98))
  fmt.Println(htb.Delete(99))
  fmt.Println(htb.Length())
}
