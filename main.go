package main

import (
	"healthcare-app/healthcare-backend/config"
	"healthcare-app/healthcare-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB() // initialize database (auto-migrate)
	router := gin.Default()

	routes.RegisterRoutes(router, db) // register all routes

	router.Run(":8080")
}
