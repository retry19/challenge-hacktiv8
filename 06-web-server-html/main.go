package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/retry19/challenge-hacktiv8/06-web-server-html/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.LoadHTMLGlob("routes/templates/*")

	routes.BuildRoutes(r)

	fmt.Println("Listening on port", port)

	r.Run(fmt.Sprintf(":%s", port))
}
