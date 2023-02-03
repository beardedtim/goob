package http

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func StartServer(serverName string) *gin.Engine {
	server := gin.Default()

	server.Use(otelgin.Middleware(serverName))
	server.Use(RequestId)
	server.Use(CustomHeaders)

	return server
}
