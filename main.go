package main

import (
	"github.com/volodymyrprokopyuk/go-ads/cc"
)


func main() {
  // * mutex
  // cc.MtxCounter()
  cc.RRWMutexPrefersReaders()
}
