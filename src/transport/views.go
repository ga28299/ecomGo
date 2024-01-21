package transport

import "github.com/gin-gonic/gin"

func ShowSignupForm(c *gin.Context) {
	// You can render an HTML template for the sign-up form here
	c.HTML(200, "signup_form.html", nil)

}

func ShowLoginForm(c *gin.Context) {
	// You can render an HTML template for the login form here
	c.HTML(200, "login_form.html", nil)
}
