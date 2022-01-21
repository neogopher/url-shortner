// Package url holds all the domain logic for URL.
package url

import (
	"errors"
	"url-shortner/pkg/domain/aggregate"
)

var (
	ErrURLNotFound  = errors.New("URL was not found")
	ErrURLAddFailed = errors.New("URL could not be added to the repository")
)

// Repository defines the set of functions that a url repository needs to implement.
type Repository interface {
	Add(aggregate.URL) error
	Get(string) (aggregate.URL, error)
}
