package main

import (
	"mckp/goob/http"
	"mckp/goob/monitoring"

	"github.com/gin-gonic/gin"
)

func onShutdown() {
	monitoring.ShutDownTracer()
}

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

/*
List of Routes that we want to attach to the HTTP Server
*/
var routes = []Route{
	{
		Name:    "Hello World",
		Path:    "/",
		Method:  "GET",
		Handler: http.HelloWorld,
	},
	{
		Name:    "Get User By Id",
		Path:    "/users/:id",
		Method:  "GET",
		Handler: http.GetUserById,
	},
}

func applyRoutes(server *gin.Engine) *gin.Engine {
	for _, route := range routes {
		/* I wish this was not a switch and just dynamic lookup but reflect is confusing */
		switch route.Method {
		case "GET":
			server.GET(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
		case "POST":
			server.POST(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
		case "PUT":
			server.PUT(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
		case "PATCH":
			server.PATCH(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
		}
	}

	return server
}

func applyMiddleware(server *gin.Engine) *gin.Engine {
	server.Use(http.TrackIncomingRequestCount)

	return server
}

func main() {
	defer onShutdown()

	name := "HTTP"

	monitoring.ConfigureEnv(name)

	server := applyRoutes(applyMiddleware(http.StartServer(name)))

	server.Run()
}
