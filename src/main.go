package main

import (
	"log"
	"os"

	"github.com/ga28299/ecomGo/db"
	"github.com/ga28299/ecomGo/transport"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	esConStr := os.Getenv("ELASTICSEARCH_CONNECTION_STRING")
	pgConStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	log.Println("Connecting to PostgreSQL with connection string:", pgConStr)
	log.Println("Connecting to Elasticsearch with connection string:", esConStr)

	_, _, err := db.SetDB(pgConStr, esConStr)
	if err != nil {
		log.Fatal("Failed to connect to dbs", err)
		return
	}

	defer db.CloseDB()

	router := gin.New()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*/**.html")
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
