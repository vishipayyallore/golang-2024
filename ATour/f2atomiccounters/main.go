package main

import (
	utl "autilities"
	"sync"
	"sync/atomic"
	"time"
)

var header = utl.Header{}

func main() {
	/*
		The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools.
		There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters
		accessed by multiple goroutines.
	*/
	header.DisplayHeader("Showing Atomic Counters")

	// We’ll use an atomic integer type to represent our (always-positive) counter.
	var ops atomic.Uint64

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// This limiter channel will receive a value every 100 milliseconds. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(100 * time.Millisecond)
	iters := 50

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < iters; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// To atomically increment the counter we use Add.
				ops.Add(1)
			}
			wg.Done()
		}()

		utl.PLine("Value of OPS : ", ops.Load(), " at i = ", i, " Time: ", time.Now())
		<-limiter
	}

	// Wait until all the goroutines are done.
	wg.Wait()

	// Here no goroutines are writing to ‘ops’, but using Load it’s safe to atomically read a value even while other goroutines are (atomically) updating it.
	utl.PLine("Value of OPS : ", ops.Load(), " at i = ", iters, " Time: ", time.Now())
}
