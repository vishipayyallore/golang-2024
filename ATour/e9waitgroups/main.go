package main

import (
	utl "autilities"
	"sync"
	"time"
)

var header = utl.Header{}

func main() {
	// To wait for multiple goroutines to finish, we can use a wait group.
	header.DisplayHeader("Showing WaitGroups")

	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()

	// Note that this approach has no straightforward way to propagate errors from workers.
	// For more advanced use cases, consider using the errgroup package.
}

// This is the function we’ll run in every goroutine. Sleep to simulate an expensive task.
func worker(id int) {
	utl.PFmted("Worker %d starting\n", id)
	time.Sleep(time.Second)
	utl.PFmted("Worker %d done\n", id)
}
