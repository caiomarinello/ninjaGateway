package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	client := &http.Client{} // Reusable HTTP client
	
	r := gin.Default()

	forwardRequest := func(c *gin.Context, url string) {
		resp, err := client.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// Copy response status, headers, and body
		c.Status(resp.StatusCode)
		for key, values := range resp.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}
		io.Copy(c.Writer, resp.Body)
	}

	r.GET("/product", func(c *gin.Context) {
		forwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product")
	})

	r.GET("/product/:productId", func(c *gin.Context) {
		forwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product/" + c.Param("productId"))
	})

	r.Run(":8085")
}
