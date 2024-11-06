package listeners

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/retry19/challenge-hacktiv8/08-assignment/controllers"
	_ "github.com/retry19/challenge-hacktiv8/08-assignment/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Hacktiv8 - 08 Assignment
// @version 0.1
// @description A simple API for Hacktiv8 - 08 Assignment
// @host localhost:8080
// @BasePath /
// @contact.name Reza Rachmanuddin
// @contact.url https://github.com/retry19
// @contact.email rezarahmanudin@gmail.com
func StartHttpListener() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	orderRouter := r.Group("/orders")
	orderRouter.GET("", controllers.GetOrders)
	orderRouter.POST("", controllers.CreateOrder)
	orderRouter.PUT(":id", controllers.UpdateOrder)
	orderRouter.DELETE(":id", controllers.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)
}
