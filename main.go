package main

import (
	"github.com/volodymyrprokopyuk/go-ads/concur"
	"github.com/volodymyrprokopyuk/go-ads/pattern"
)

func patternTest() {
  // pattern.PatDecorator()
  pattern.PatOptions()
}

func concurTest() {
  // * mutex
  // concur.MtxCounter()
  // concur.RRWMutexPrefersReaders()
  // concur.RWWMutexPrefersWriters()

  // * condition
  // concur.CndBalance()
  // concur.CndAllJoined()

  // * semaphore
  // concur.SemCndConcurLimit()
  // concur.SemChConcurLimit()

  // * wait group
  // concur.WGAllDone()

  // * barrier
  // concur.BarSyncRounds()

  // * channel
  // concur.ChSyncAsyncPipe()
  // concur.ChMultiSendReceive()
  // concur.ChEarlyExit()
  // concur.ChFanOutFanIn()
  // concur.ChFanOutFanIn2()
  // concur.ChBroadcast()
  // concur.ChPipeline()
  // concur.ChErrorHandling()
  // concur.ChTee()
  // concur.ChMerge()
  // concur.ChHeartbeat()
  // concur.ChAsyncRateLimiter()
  // concur.ChPool()

  // * context
  // concur.CtxCancelTimeout()
  // concur.CtxGracefulTermination()
}

func concurProblem() {
  concur.ChSieveOfEratosthenes()
}


func main() {
  patternTest()
  // concurTest()
  // concurProblem()
}
