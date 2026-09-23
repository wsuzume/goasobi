package pkg

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func EchoHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
		"query":   c.Request.URL.Query(),
		"headers": c.Request.Header,
		"body":    string(body),
	})
}

// NewFrontendHandler serves the static assets built into dir (see
// goasobi-frontend and goasobi/Dockerfile). Requests for a path that exists
// as a file under dir are served as-is; any other path falls back to
// dir/index.html so client-side routing keeps working.
func NewFrontendHandler(dir string) gin.HandlerFunc {
	fileServer := http.FileServer(http.Dir(dir))

	return func(c *gin.Context) {
		// c.Request.URL.Path is always rooted ("/..."), so Clean can never
		// escape dir via "..".
		requestPath := filepath.Clean(c.Request.URL.Path)

		if info, err := os.Stat(filepath.Join(dir, requestPath)); err == nil && !info.IsDir() {
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		c.File(filepath.Join(dir, "index.html"))
	}
}
