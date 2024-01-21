package transport

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/users/signup", ShowSignupForm)
	incomingRoutes.POST("/users/signup", Signup)

	incomingRoutes.POST("/users/login", Login)
	incomingRoutes.GET("/users/login", ShowLoginForm)

	incomingRoutes.POST("/admin/addproduct", ProductViewerAdmin)
	incomingRoutes.GET("/users/productview", SearchProduct)
	incomingRoutes.POST("/users/search", SearchProductQuery)
}
