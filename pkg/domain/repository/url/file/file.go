// Package memory is a file-based implementation of the URL repository.
package file

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"url-shortner/pkg/domain/aggregate"
	"url-shortner/pkg/domain/repository/url"
)

// FileRepository implements the URL repository.
type FileRepository struct {
	urls     map[string]aggregate.URL
	mu       sync.Mutex
	filename string
}

// NewFileRepository is a factory function for creating FileRepository.
func NewFileRepository(filename string) *FileRepository {
	return &FileRepository{
		urls:     make(map[string]aggregate.URL),
		filename: filename,
	}
}

// Add adds URL to repository.
func (fr *FileRepository) Add(url aggregate.URL) error {
	_, ok := fr.urls[url.GetID()]
	if !ok {

		fr.mu.Lock()
		fr.urls[url.GetID()] = url

		// Additionaly write new links to file.
		f, err := os.OpenFile(fr.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		fmt.Fprintln(f, strings.Join([]string{url.GetID(), url.GetFullPath()}, ","))
		f.Close()

		fr.mu.Unlock()
	}

	return nil
}

// Get gets URL from the repository.
func (fr *FileRepository) Get(shortCode string) (aggregate.URL, error) {
	u, ok := fr.urls[shortCode]
	if !ok {
		return aggregate.URL{}, url.ErrURLNotFound
	}

	return u, nil
}
