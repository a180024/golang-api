package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	router := SetupRouter()

	body := map[string]interface{}{
		"UserName": "test",
		"Password": "test",
		"Email":    "test@example.com",
	}
	bodyJson, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	req.Body.Close()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"User successfully registered!\"}", w.Body.String())
}

func TestLogin(t *testing.T) {
	router := SetupRouter()

	body := map[string]interface{}{
		"UserName": "test",
		"Password": "test",
	}
	bodyJson, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	req.Body.Close()

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Token")
}
