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
  // cc.CndAllJoined()

  // * semaphore
  // cc.SemConcurrencyLimit()

  // * wait group
  // cc.WGAllDone()

  // * barrier
  // cc.BarSyncRounds()

  // * channel
  cc.ChSyncAsyncPipe()
}


func main() {
  concurrency()
}
