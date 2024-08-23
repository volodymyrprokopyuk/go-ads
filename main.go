package main

import (
	"github.com/volodymyrprokopyuk/go-ads/cc"
	"github.com/volodymyrprokopyuk/go-ads/cc/prb"
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
  // cc.ChSyncAsyncPipe()
  // cc.ChEarlyExist()
  // cc.ChFanOutFanIn()
  // cc.ChBroadcast()
  // cc.ChPipeline()
  // cc.ChErrorHandling()
  // cc.ChTee()
  cc.ChMerge()
  // cc.ChHeartbeat()
  // cc.ChAsyncRateLimiter()

  // * context
  // cc.CtxCancelTimeout()
  // cc.CtxGracefulTermination()
}

func ccProblem() {
  prb.ChSieveOfEratosthenes()
}


func main() {
  concurrency()
  // ccProblem()
}
