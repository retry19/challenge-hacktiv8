package listeners

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/retry19/challenge-hacktiv8/10-assignment/controllers"
)

func StartHttpListener() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.POST("/auto-reload", controllers.AutoReloadController)

	r.Run(":" + port)
}
