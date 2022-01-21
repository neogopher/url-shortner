// Package app is the starting point of the application.
package app

import (
	"url-shortner/pkg/domain/repository/url/memory"
	"url-shortner/pkg/handlers"
	"url-shortner/pkg/services"

	"github.com/gin-gonic/gin"
)

func Start() {
	repo := memory.NewMemoryRepository()
	service := services.NewShortnerService(repo)
	handler := handlers.NewHTTPHandler(*service)

	r := gin.Default()

	r.POST("/shorten", handler.GetShortenedURL)

	r.Run(":8080")
}
