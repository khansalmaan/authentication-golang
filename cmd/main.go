package main

import (
	"github.com/gin-gonic/gin"

	database "go-rest/api/external/db"
	routes "go-rest/api/routes"
	config "go-rest/config"
)

func main() {
	r := gin.Default()
	// Initialize the Config
	config.Init()
	// Initialize database
	database.Init()
	// Initialize the routes
	r = routes.Init(r)
	r.Run(":" + config.Config.Port)
}
