package pkg

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestHelloHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	HelloHandler(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if body["message"] != "Hello, World!" {
		t.Errorf("expected message %q, got %q", "Hello, World!", body["message"])
	}
}

func TestEchoHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/echo?foo=bar", strings.NewReader("hello"))
	c.Request.Header.Set("X-Test", "value")

	EchoHandler(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var body struct {
		Method  string              `json:"method"`
		Path    string              `json:"path"`
		Query   map[string][]string `json:"query"`
		Headers map[string][]string `json:"headers"`
		Body    string              `json:"body"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	if body.Method != http.MethodPost {
		t.Errorf("expected method %q, got %q", http.MethodPost, body.Method)
	}
	if body.Path != "/echo" {
		t.Errorf("expected path %q, got %q", "/echo", body.Path)
	}
	if got := body.Query["foo"]; len(got) != 1 || got[0] != "bar" {
		t.Errorf("expected query foo=bar, got %v", body.Query["foo"])
	}
	if got := body.Headers["X-Test"]; len(got) != 1 || got[0] != "value" {
		t.Errorf("expected header X-Test=value, got %v", body.Headers["X-Test"])
	}
	if body.Body != "hello" {
		t.Errorf("expected body %q, got %q", "hello", body.Body)
	}
}
