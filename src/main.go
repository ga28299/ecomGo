package main

import (
	"log"
	"os"

	"github.com/ga28299/ecomGo/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	transport.UserRoutes(router.Group("/api/"))
	// router.Use(middleware.Authenticattion())

	// app := transport.NewApp(db.ProductData(db.Client, "Products"), db.UserData(db.CLient), "Users")

	// router.GET("/addtocart", app.AddToCart())
	// router.GET("/removeitem", app.RemoveItem())
	// router.GET("/cartcheckout", app.CartBuy())
	// router.GET("/instantsale", app.InstantSale())

	log.Fatal(router.Run(":" + port))

}
