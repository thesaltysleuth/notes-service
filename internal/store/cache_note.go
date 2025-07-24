package store

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thesaltysleuth/notes-service/internal/models"
)

// CachedNoteStore wraps another NoteStore and adds Redis caching
type CachedNoteStore struct {
	backend NoteStore
	rdb *redis.Client
	ttl time.Duration
}

func NewCachedNoteStore(backend NoteStore, redisAddr string, ttl time.Duration) *CachedNoteStore {
	rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
	return &CachedNoteStore{backend: backend, rdb: rdb, ttl:ttl}
}

func cacheKey(owner string) string { return fmt.Sprintf("notes:%s",owner) }

func (c *CachedNoteStore) Add(owner, title, content string) models.Note {
	n := c.backend.Add(owner, title, content)
	// invalidate cache
	_ = c.rdb.Del(context.Background(), cacheKey(owner)).Err()
	return n
}

func (c *CachedNoteStore) GetAll(owner string) []models.Note {
	ctx := context.Background()
	key := cacheKey(owner)

	// 1. Try cache
	if raw, err := c.rdb.Get(ctx, key).Bytes(); err == nil {
		var cached []models.Note
		if json.Unmarshal(raw, &cached) == nil {
			return cached
		}
	}

	// 2. Fallback to backend
	res := c.backend.GetAll(owner)


	// 3. Set cache asynchronously (best-effort)
	go func(data []models.Note) {
		if enc,err := json.Marshal(data); err == nil {
			_ = c.rdb.Set(ctx, key, enc, c.ttl).Err()
		}
	}(res)

	return res
}
