// Package services contains all services that work with Handlers and repositories to perform an action.
package services

import "url-shortner/pkg/domain/repository/url"

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
