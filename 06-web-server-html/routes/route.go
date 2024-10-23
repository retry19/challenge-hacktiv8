package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/retry19/challenge-hacktiv8/06-web-server-html/controllers"
)

func BuildRoutes(r *gin.Engine) {
	home := r.Group("/")
	home.GET("/", controllers.ShowLoginPage)
	home.POST("/", controllers.SubmitLogin)
}
