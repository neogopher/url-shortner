// Package handlers contains all handlers that call services
package handlers

import (
	"url-shortner/pkg/services"

	"github.com/gin-gonic/gin"
)

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

// GetShortenedURL accepts a URL as an argument over POST and returns a shortened URL.
func (h *HTTPHandler) GetShortenedURL(c *gin.Context) {

}

// GetLongURL accepts a shortened URL as an argument over POST and returns the original full URL.
func (h *HTTPHandler) GetLongURL(c *gin.Context) {

}
