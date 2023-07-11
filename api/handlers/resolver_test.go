package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	handler := &Handler{engine: router}
	handler.setHandlers()
	return router
}

func TestResolveShorten(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := SetupRouter()
	t.Run("valid URL", func(t *testing.T) {
		reqBody := struct {
			URL string `json:"url"`
		}{
			URL: "https://example.com",
		}
		reqJSON, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/shorten", bytes.NewReader(reqJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusAccepted, rec.Code)
		expectedBody := `{"message":"Shorten"}`
		assert.JSONEq(t, expectedBody, rec.Body.String())
	})

	t.Run("no URL", func(t *testing.T) {
		// Create a request with a valid URL
		reqBody := struct {
			URL string `json:"url"`
		}{}
		reqJSON, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/shorten", bytes.NewReader(reqJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		expectedBody := `{"error":"Key: 'request.URL' Error:Field validation for 'URL' failed on the 'required' tag"}`
		assert.JSONEq(t, expectedBody, rec.Body.String())
	})

	t.Run("invalid URL", func(t *testing.T) {
		reqBody := struct {
			URL string `json:"url"`
		}{
			URL: "abcd123",
		}
		reqJSON, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/shorten", bytes.NewReader(reqJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		expectedBody := `{"invalid URL": "abcd123"}`
		assert.JSONEq(t, expectedBody, rec.Body.String())
	})
}
