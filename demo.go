package skywalkingplugin

import (
	"context"
	"fmt"
	"github.com/rs/xid"
	"net/http"
	"time"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

// Demo a Demo plugin.
type Demo struct {
	next    http.Handler
	headers map[string]string
	name    string
}

// New created a new Demo plugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Demo{
		headers: config.Headers,
		next:    next,
		name:    name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	before := time.Now()
	a.next.ServeHTTP(rw, req)
	fmt.Printf("traceId: %s, 请求URL: %s, 耗时: %s\n", xid.New().String(), req.URL.Path, time.Since(before))
}
