package main

import (
	"fmt"
	"runtime"

	"github.com/thesaltysleuth/notes-service/internal/worker"
)

func main() {
	N := runtime.NumCPU()
	jobs := make(chan worker.Job)
	go func() {
		for i := 1; i <=700; i++ {
			jobs <- worker.Job(i)
		}
		close(jobs)
	}()

	for r := range worker.StartPool(N, jobs) {
		fmt.Println(r)
	}
}
