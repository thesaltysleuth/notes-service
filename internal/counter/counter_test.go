package counter

import ( 
	"testing"
)

type incVal interface {
	Inc()
	Value() int64
}

func benchCounter(b *testing.B, c incVal) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}


func BenchmarkMutexCounter(b *testing.B) {
	benchCounter(b, &MutexCounter{})
}

func BenchmarkAtomicCounter(b *testing.B) {
	benchCounter(b, &AtomicCounter{})
}

