// Package aggregate holds aggregates that combine multiple entities into full objects.
package aggregate

import (
	"errors"
	"url-shortner/internal/hash"
	"url-shortner/pkg/domain/entity"

	"net/url"
)

var (
	ErrInvalidPath = errors.New("path is invalid")
	ErrEmptyPath   = errors.New("path should not be empty")
)

// URL represents a type of Link.
type URL struct {
	link *entity.Link
}

// NewURL is a factory function to create URL.
// It will validate the supplied path.
func NewURL(path string) (URL, error) {
	if path == "" {
		return URL{}, ErrEmptyPath
	}

	if _, err := url.ParseRequestURI(path); err != nil {
		return URL{}, ErrInvalidPath
	}

	shortCode := hash.GenerateShortCode(path)

	return URL{
		link: &entity.Link{
			ID:       shortCode,
			FullPath: path,
		},
	}, nil
}
