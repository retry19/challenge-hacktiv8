package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/retry19/challenge-hacktiv8/06-web-server-html/services"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func SubmitLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := services.FindUserByEmail(email)
	if user == nil || user.Password != password {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": "Invalid email or password",
			"email": email,
		})
		return
	}

	c.HTML(http.StatusOK, "home.html", user)
}
