package pkg

import (
	"io"
	"net/http"

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
