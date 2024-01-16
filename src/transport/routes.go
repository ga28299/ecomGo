so in my context in my routes.go file :package transport

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/users/signup", Signup)
	incomingRoutes.POST("/users/signup", Signup)
	
	incomingRoutes.POST("/users/login", Login)
	incomingRoutes.POST("/admin/addproduct", ProductViewerAdmin)
	incomingRoutes.GET("/users/productview", SearchProduct)
	incomingRoutes.POST("/users/search", SearchProductQuery)
}
