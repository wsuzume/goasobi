package pkg

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewProxyServer(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Backend-Path", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer backend.Close()

	target, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatalf("failed to parse backend URL: %v", err)
	}

	// A real listener is used, rather than httptest.NewRecorder, because
	// httputil.ReverseProxy relies on http.CloseNotifier support that
	// ResponseRecorder does not provide.
	proxy := httptest.NewServer(NewProxyServer(target))
	defer proxy.Close()

	resp, err := http.Get(proxy.URL + "/some/path?foo=bar")
	if err != nil {
		t.Fatalf("request to proxy failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if got := resp.Header.Get("X-Backend-Path"); got != "/some/path" {
		t.Errorf("expected backend to receive path %q, got %q", "/some/path", got)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if string(body) != "ok" {
		t.Errorf("expected body %q, got %q", "ok", string(body))
	}
}
