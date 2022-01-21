// Package memory is an in-memory implementation of the URL repository.
package memory

import (
	"sync"
	"url-shortner/pkg/domain/aggregate"
	"url-shortner/pkg/domain/repository/url"
)

// MemoryRepository implements the URL repository.
type MemoryRepository struct {
	urls map[string]aggregate.URL
	mu   sync.Mutex
}

// NewRepository is a factory function for creating MemoryRepository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		urls: make(map[string]aggregate.URL),
	}
}

func (mr *MemoryRepository) Add(url aggregate.URL) error {
	mr.mu.Lock()
	mr.urls[url.GetID()] = url
	mr.mu.Unlock()

	return nil
}

func (mr *MemoryRepository) Get(shortCode string) (aggregate.URL, error) {
	u, ok := mr.urls[shortCode]
	if !ok {
		return aggregate.URL{}, url.ErrURLNotFound
	}

	return u, nil
}
