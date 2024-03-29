package router

import (
	"github.com/bxmon/diapi-mock-server/controller"
	"github.com/gin-gonic/gin"
)

// AppRoute defines application's route structure
type AppRoute struct {
	Group            string
	Protected        bool
	Routes           []Route
	GroupMiddlewares []gin.HandlerFunc
}

// Route defines a single route, e.g. a human readable name, HTTP method and
// the pattern, the function that will execute when the route is called
type Route struct {
	Method           string
	Pattern          string
	RouteMiddlewares []gin.HandlerFunc
}

// NewRouters initializes all scraper routers
func NewRouters(c *controller.Controller, e *gin.Engine) {
	ar := GetRoutes(c)
	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, rs := range ar {
		groupRoute := e.Group(rs.Group, rs.GroupMiddlewares...)

		for _, r := range rs.Routes {
			switch r.Method {
			case "GET":
				groupRoute.GET(r.Pattern, r.RouteMiddlewares...)
			case "POST":
				groupRoute.POST(r.Pattern, r.RouteMiddlewares...)
			case "PUT":
				groupRoute.PUT(r.Pattern, r.RouteMiddlewares...)
			case "PATCH":
				groupRoute.PATCH(r.Pattern, r.RouteMiddlewares...)
			case "HEAD":
				groupRoute.HEAD(r.Pattern, r.RouteMiddlewares...)
			case "OPTIONS":
				groupRoute.OPTIONS(r.Pattern, r.RouteMiddlewares...)
			case "DELETE":
				groupRoute.DELETE(r.Pattern, r.RouteMiddlewares...)
			}
		}
	}
}

// GetRoutes initializes all account routers
func GetRoutes(c *controller.Controller) []AppRoute {
	return []AppRoute{
		AppRoute{
			Group:            "/api/v1",
			GroupMiddlewares: []gin.HandlerFunc{},
			Routes: []Route{
				Route{
					Method:  "POST",
					Pattern: "/users/register",
					RouteMiddlewares: []gin.HandlerFunc{
						c.AddUserHandler,
					},
				},
				Route{
					Method:  "GET",
					Pattern: "/users/:userid",
					RouteMiddlewares: []gin.HandlerFunc{
						c.GetUserByIDHandler,
					},
				},
				Route{
					Method:  "PATCH",
					Pattern: "/users/:userid",
					RouteMiddlewares: []gin.HandlerFunc{
						c.UpdateUserHandler,
					},
				},
				Route{
					Method:  "DELETE",
					Pattern: "/users/:userid",
					RouteMiddlewares: []gin.HandlerFunc{
						c.DeleteUserByIDHandler,
					},
				},
				Route{
					Method:  "GET",
					Pattern: "/users",
					RouteMiddlewares: []gin.HandlerFunc{
						c.GetUsersHandler,
					},
				},
			},
		},
	}
}
