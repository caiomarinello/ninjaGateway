package main

import (
	"log"
	"os"

	"github.com/caiomarinello/ninjaGateway/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Check if the UPSTREAM_URL environment variable is already set
	if _, exists := os.LookupEnv("UPSTREAM_URL"); !exists {
		// If UPSTREAM_URL is not set, load environment variables from the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/products")
	})
	r.GET("/product/:productId", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product/" + c.Param("productId"))
	})
	r.POST("/product", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product")
	})
	r.PUT("/product/:productId", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product/" + c.Param("productId"))
	})
	r.DELETE("/product/:productId", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product/" + c.Param("productId"))
	})

	r.Run(":8085")
}
