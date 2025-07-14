package buffer

import "testing"

//naive: let append grow capacity on demand
func BenchmarkAppendNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf Buffer
		for j := 0; j < 1_000; j++ {
			buf.Write([]byte{'x'})
		}
	}
}


//pre-allocated: reserve exactly what we need
func BenchmarkAppendPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := Buffer{data: make([]byte, 0, 1_000)}
		for j := 0; j < 1_000; j++ {
			buf.Write([]byte{'x'})
		}
	}
}

