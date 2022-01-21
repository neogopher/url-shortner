// Package memory is a file-based implementation of the URL repository.
package file

import (
	"sync"
	"url-shortner/pkg/domain/aggregate"
)

// FileRepository implements the URL repository.
type FileRepository struct {
	urls     map[string]aggregate.URL
	mu       sync.Mutex
	filename string
}

// NewFileRepository is a factory function for creating FileRepository.
func NewFileRepository() *FileRepository {
	return &FileRepository{
		urls: make(map[string]aggregate.URL),
	}
}

// Add adds URL to repository.
func (fr *FileRepository) Add(url aggregate.URL) error {
	return nil
}

// Get gets URL from the repository.
func (fr *FileRepository) Get(shortCode string) (aggregate.URL, error) {
	return aggregate.URL{}, nil
}
