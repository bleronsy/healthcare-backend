package main

import (
	"fmt"
	"time"

	"healthcare-app/healthcare-backend/config"
	"healthcare-app/healthcare-backend/routes"

	_ "healthcare-app/healthcare-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDatabase()

	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s - [%s] \"%s %s\" %d %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	r.Use(gin.Recovery())

	routes.SetupRoutes(r)
	r.Run(":8080")
}
