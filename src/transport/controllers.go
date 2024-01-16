package transport

import "github.com/gin-gonic/gin"

func Signup(c *gin.Context) {
	
	
}

func Login(c *gin.Context) {

}

func ProductViewerAdmin(c *gin.Context) {

}

func SearchProduct(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", []byte("Hello World"))

}

func SearchProductQuery(c *gin.Context) {
	c.JSON(200, gin.H{})
}
