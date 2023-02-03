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

type RouteGroup struct {
	Name   string
	Prefix string
	Routes []Route
}

/*
List of Routes that we want to attach to the HTTP Server
*/
var rootRoutes = RouteGroup{
	Name:   "Root Routes",
	Prefix: "",
	Routes: []Route{
		{
			Name:    "Hello World",
			Path:    "/",
			Method:  "GET",
			Handler: http.HelloWorld,
		},
	},
}

var userRoutes = RouteGroup{
	Name:   "User Routes",
	Prefix: "/users",
	Routes: []Route{
		{
			Name:    "Get User By Id",
			Path:    "/:id",
			Method:  "GET",
			Handler: http.GetUserById,
		},
	},
}

var routes = []RouteGroup{
	rootRoutes,
	userRoutes,
}

func applyRoutes(server *gin.Engine) *gin.Engine {
	for _, routeGroup := range routes {
		group := server.Group(routeGroup.Prefix)

		for _, route := range routeGroup.Routes {
			/* I wish this was not a switch and just dynamic lookup but reflect is confusing */
			switch route.Method {
			case "GET":
				group.GET(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
			case "POST":
				group.POST(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
			case "PUT":
				group.PUT(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
			case "PATCH":
				group.PATCH(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
			case "DELETE":
				group.DELETE(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
			case "HEAD":
				group.HEAD(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
			case "OPTIONS":
				group.OPTIONS(route.Path, monitoring.WrapMiddleware(route.Handler, route.Name))
			}
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
