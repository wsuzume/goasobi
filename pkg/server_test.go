package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServerRoutes(t *testing.T) {
	r := NewServer()

	tests := []struct {
		name   string
		method string
		path   string
		want   int
	}{
		{"root", http.MethodGet, "/", http.StatusOK},
		{"echo get", http.MethodGet, "/echo", http.StatusOK},
		{"echo post", http.MethodPost, "/echo", http.StatusOK},
		{"unknown", http.MethodGet, "/does-not-exist", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, tt.path, nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.want {
				t.Errorf("%s %s: expected status %d, got %d", tt.method, tt.path, tt.want, w.Code)
			}
		})
	}
}
