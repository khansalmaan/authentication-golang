package routes

import (
	"github.com/gin-gonic/gin"

	routesV1 "go-rest/api/routes/v1"
)

// Init initializes the routes for the application
func Init(r *gin.Engine) *gin.Engine {

	// base route
	r.GET("/", gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go REST API",
		})
	}))

	// Initialize the routes
	v1 := r.Group("/api/v1")

	// Initialize the user routes
	usersRoutes := v1.Group("/users")
	routesV1.UserRoutes(usersRoutes)

	// v2 routes
	// v2 := r.Group("/api/v2")
	//  Initialize the user routes v2
	// usersRoutesV2 := v2.Group("/users")
	// routesV2.UsersRoutes(usersRoutesV2)

	return r
}
