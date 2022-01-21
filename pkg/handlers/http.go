// Package handlers contains all handlers that call services
package handlers

import "url-shortner/pkg/services"

// HTTPHandler interacts with ShortnerService.
type HTTPHandler struct {
	service services.ShortnerService
}

// NewHTTPHandler is a factory function for creating HTTPHandler.
func NewHTTPHandler(s services.ShortnerService) *HTTPHandler {
	return &HTTPHandler{
		service: s,
	}
}
