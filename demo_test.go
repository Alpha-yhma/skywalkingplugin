package skywalkingplugin_test

import (
	"context"
	"github.com/traefik/skywalkingplugin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemo(t *testing.T) {
	cfg := skywalkingplugin.CreateConfig()
	cfg.Headers["X-Host"] = "[[.Host]]"
	cfg.Headers["X-Method"] = "[[.Method]]"
	cfg.Headers["X-URL"] = "[[.URL]]"
	cfg.Headers["X-URL"] = "[[.URL]]"
	cfg.Headers["X-Demo"] = "test"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := skywalkingplugin.New(ctx, next, cfg, "demo-plugin", t)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/xxxx", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Host", "localhost")
	assertHeader(t, req, "X-URL", "http://localhost/xxxx")
	assertHeader(t, req, "X-Method", "GET")
	assertHeader(t, req, "X-Demo", "test")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
