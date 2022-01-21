// Package handlers contains all handlers that call services
package handlers

import (
	"fmt"
	"net/http"
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
	var req struct {
		URL string `json:"url" binding:"required"`
	}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	stub, err := h.service.Shorten(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	shortURL := fmt.Sprintf("%s://%s/%s", scheme, c.Request.Host, stub)

	type response struct {
		ShortURL string `json:"shortUrl"`
	}

	resp := response{
		ShortURL: shortURL,
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "shortened URL generated", "data": resp})
}

// GetLongURL accepts a shortened URL as an argument over GET and returns the original full URL.
func (h *HTTPHandler) GetLongURL(c *gin.Context) {

}
