// Package memory is an in-memory implementation of the URL repository.
package memory

import (
	"sync"
	"url-shortner/pkg/domain/aggregate"
)

// MemoryRepository implements the URL repository.
type MemoryRepository struct {
	urls map[string]aggregate.URL
	sync.Mutex
}

// NewRepository is a factory function for creating MemoryRepository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		urls: make(map[string]aggregate.URL),
	}
}

func (mr *MemoryRepository) Add(aggregate.URL) error {
	return nil
}

func (mr *MemoryRepository) Get(shortCode string) (aggregate.URL, error) {
	return aggregate.URL{}, nil
}
