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

type Int int

func (i Int) String() string {
  return strconv.Itoa(int(i))
}

func main() {
  // collision()

  htb := ads.NewHTable[Int, string](101)

  htb.Set(Int(66), ">66")
  htb.Set(Int(98), ">98")
  htb.Set(Int(98), ">>98")

  fmt.Println(htb.Length())
  for k, v := range htb.Entries() {
    fmt.Println(k, v)
  }

  fmt.Println(htb.Get(Int(66)))
  fmt.Println(htb.Get(Int(98)))
  fmt.Println(htb.Get(Int(99)))

  fmt.Println(htb.Delete(Int(66)))
  fmt.Println(htb.Delete(Int(98)))
  fmt.Println(htb.Delete(Int(99)))
  fmt.Println(htb.Length())

  fmt.Println(2 ^ 3)
}
