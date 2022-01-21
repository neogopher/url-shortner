// Package handlers contains all handlers that call services
package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortner/pkg/domain/repository/url/memory"
	"url-shortner/pkg/services"

	"github.com/gin-gonic/gin"
)

func TestHTTPHandler_GetShortenedURL(t *testing.T) {
	path := "http://www.google.com"
	// _ := hash.GenerateShortCode(path)

	repo := memory.NewMemoryRepository()
	service := services.NewShortnerService(repo)
	h := NewHTTPHandler(*service)

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/shorten", h.GetShortenedURL)

	tests := []struct {
		name        string
		requestBody []byte
		wantCode    int
	}{
		{
			name:        "url argument present",
			requestBody: []byte(`{"url": "` + path + `" }`),
			wantCode:    http.StatusOK,
		},
		{
			name:        "url argument absent",
			requestBody: []byte(`{}`),
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "url invalid",
			requestBody: []byte(`{"url": "invalid url"}`),
			wantCode:    http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("POST /shorten Want: %d Got: %d", http.StatusOK, w.Code)
			}
		})
	}

}
