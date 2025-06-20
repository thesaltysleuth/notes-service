package worker

import (
	"sync"
	"time"
//	"fmt"
)

type Job int  		// for today, an int we'll square
type Result int 	// squared value


func StartPool(concurrency int, jobs <-chan Job) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				time.Sleep(100 * time.Millisecond) //simulate heavy task
				res := Result(j*j)
				// fmt.Printf("Worker %d processed %d -> %d \n", id, j, res)
				results <- res
			}
		}(i)
	}

	go func() {
		wg.Wait();
		close(results) 
	}()
	return results
}


