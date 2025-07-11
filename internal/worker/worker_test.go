package worker_test

import (
	"runtime"
	"testing"

	"github.com/thesaltysleuth/notes-service/internal/worker"
)

func BenchmarkPoolWaitGroup(b *testing.B) {
	jobs := make(chan worker.Job, b.N)
	for i := 0; i< b.N; i++ {
		jobs <- worker.Job(i)
	}
	close(jobs)

	for range worker.StartPool(runtime.NumCPU(), jobs) {
		//drain results
	}
}

func BenchmarkSerial(b *testing.B) {
	for i := 0; i< b.N; i++ {
		_ = i * i
	}
}

func BenchmarkPoolChanOnly(b *testing.B) {
	jobs := make(chan worker.Job, b.N)
	for i := 0; i< b.N; i++ {
		jobs <- worker.Job(i)
	}
	close(jobs)

	for range worker.StartPoolChanOnly(runtime.NumCPU(), jobs) {
	}
}
