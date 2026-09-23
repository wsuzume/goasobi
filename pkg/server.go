package pkg

import "github.com/gin-gonic/gin"

// DefaultFrontendDir is where goasobi/Dockerfile places the built
// goasobi-frontend assets in the runner image.
const DefaultFrontendDir = "/usr/share/goasobi/static"

func NewServer() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", HelloHandler)
	r.Any("/echo", EchoHandler)
	r.NoRoute(NewFrontendHandler(DefaultFrontendDir))

	return r
}
