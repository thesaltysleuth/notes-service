package worker

import (
	"sync"
)

type Job int  		// for today, an int we'll square
type Result int 	// squared value


func StartPool(concurrency int, jobs <-chan Job) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				results <- Result(j*j)
			}
		}()
	}

	go func() {
		wg.Wait();
		close(results) 
	}()
	return results
}


