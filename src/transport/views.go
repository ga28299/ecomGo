package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowSignupForm(c *gin.Context) {
	// You can render an HTML template for the sign-up form here
	c.HTML(200, "signup_form.html", nil)

}

func ShowLoginForm(c *gin.Context) {
	// You can render an HTML template for the login form here
	c.HTML(200, "login_form.html", nil)
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Your E-commerce Website",
	})
}

func Nav(c *gin.Context) {
	c.HTML(http.StatusOK, "navbar.html", gin.H{})
}

func Footer(c *gin.Context) {
	c.HTML(http.StatusOK, "footer.html", gin.H{})
}

