package store

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
)

func BenchmarkGetAll(b *testing.B) {
	// Fake Redis
	mr, _ := miniredis.Run()
	defer mr.Close()

	mem := NewMemNoteStore()
	// preload 50 notes
	for i := 0; i < 50; i++ {
		mem.Add("alice", "t", "c")
	}
	cached := NewCachedNoteStore(mem, mr.Addr(), 5*time.Minute)

	// first call to warm cache
	_ = cached.GetAll("alice")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cached.GetAll("alice")
	}
}


func BenchmarkGetAllNoCache(b *testing.B) {
	mem := NewMemNoteStore()
	for i := 0; i < 50; i++ { mem.Add("alice","t","c") }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mem.GetAll("alice")
	}
}


func BenchmarkGetAll_Postgres(b *testing.B) {
    // Ensure docker‑compose is up so PG is reachable
    dsn := "postgres://notesuser:notespass@localhost:5432/notes?sslmode=disable"
    pg, _ := NewPGNoteStore(dsn)

    // preload 50 notes for user alice
    for i := 0; i < 50; i++ { pg.Add("alice", "t", "c") }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = pg.GetAll("alice")
    }
}

func BenchmarkGetAll_Postgres_Cache(b *testing.B) {
    dsn := "postgres://notesuser:notespass@localhost:5432/notes?sslmode=disable"
    pg, _ := NewPGNoteStore(dsn)

    mr, _ := miniredis.Run()                // or real Redis container
    cached := NewCachedNoteStore(pg, mr.Addr(), 5*time.Minute)

    for i := 0; i < 50; i++ { pg.Add("alice", "t", "c") }
    _ = cached.GetAll("alice")              // warm cache

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = cached.GetAll("alice")
    }
}
