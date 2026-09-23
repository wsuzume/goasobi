package pkg

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// NewProxyServer returns a gin.Engine that reverse-proxies every request it
// receives to target.
func NewProxyServer(target *url.URL) *gin.Engine {
	p := httputil.NewSingleHostReverseProxy(target)
	p.ModifyResponse = LoggingModifyResponse

	r := gin.Default()
	r.Use(LoggingMiddleware)
	r.NoRoute(gin.WrapH(p))

	return r
}

// LoggingMiddleware is an example gin middleware that logs each incoming
// request before it is forwarded to the proxy.
func LoggingMiddleware(c *gin.Context) {
	log.Printf("proxy request: %s %s", c.Request.Method, c.Request.URL.String())
	c.Next()
}

// LoggingModifyResponse is an example httputil.ReverseProxy.ModifyResponse
// hook that logs the response returned by the target server. Assign it to
// a ReverseProxy's ModifyResponse field to enable it.
func LoggingModifyResponse(resp *http.Response) error {
	log.Printf("proxy response: %s %s -> %d", resp.Request.Method, resp.Request.URL.String(), resp.StatusCode)
	return nil
}
