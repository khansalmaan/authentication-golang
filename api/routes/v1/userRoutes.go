package routesv1

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	// Initialize the routes
	r.GET("/login", gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET /api/v1/users",
		})
	}))

	return r
}
