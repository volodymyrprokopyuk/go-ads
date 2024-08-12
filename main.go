package main

import (
	"github.com/volodymyrprokopyuk/go-ads/cc"
)

func concurrency() {
  // * mutex
  // cc.MtxCounter()
  // cc.RRWMutexPrefersReaders()
  // cc.RWWMutexPrefersWriters()

  // * condition
  // cc.CndBalance()
  cc.CndAllJoined()
}


func main() {
  concurrency()
}
