// Package services contains all services that work with Handlers and repositories to perform an action.
package services

import (
	"url-shortner/pkg/domain/aggregate"
	"url-shortner/pkg/domain/repository/url"
)

// ShortnerService shortens given URL and stores it in repository.
type ShortnerService struct {
	urls url.Repository
}

// NewShortnerService is a factory function for creating ShortnerService.
func NewShortnerService(u url.Repository) *ShortnerService {
	return &ShortnerService{
		urls: u,
	}
}

func (s *ShortnerService) Shorten(fullPath string) (string, error) {
	url, err := aggregate.NewURL(fullPath)
	if err != nil {
		return "", err
	}

	s.urls.Add(url)

	return url.GetID(), nil
}

func (s *ShortnerService) GetFullPath(shortCode string) (string, error) {
	url, err := s.urls.Get(shortCode)
	if err != nil {
		return "", err
	}

	return url.GetFullPath(), nil
}
