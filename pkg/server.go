package pkg

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", HelloHandler)
	r.Any("/echo", EchoHandler)

	return r
}
