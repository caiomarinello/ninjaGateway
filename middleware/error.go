package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware function to handle all function handlers errors.
// It logs the error and sets a JSON response according to the http status code.
func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			httpResponseStatus := c.Writer.Status()
			log.Printf("Error handling request:\n"+
				"\t Code: %v\n"+
				"\t Method: %v %v\n"+
				"\t Msg: %v\n"+ //Custom message
				"\t Error: %v",
				httpResponseStatus, c.Request.Method, c.Request.URL.Path, ginErr.Meta.(map[string]interface{})["log"], ginErr)

			switch httpResponseStatus {
			case http.StatusBadRequest:
				errorMeta, ok := ginErr.Meta.(map[string]interface{})
				if ok {
					errorJsonResponse := map[string]interface{}{
						"error": errorMeta["error"],
					}
					c.JSON(http.StatusBadRequest, errorJsonResponse)
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
				}

			case http.StatusInternalServerError:
				errorMeta, ok := ginErr.Meta.(map[string]interface{})
				if ok {
					errorJsonResponse := map[string]interface{}{
						"error": errorMeta["error"],
					}
					c.JSON(http.StatusInternalServerError, errorJsonResponse)
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				}

			case http.StatusUnauthorized:
				errorMeta, ok := ginErr.Meta.(map[string]interface{})
				if ok {
					errorJsonResponse := map[string]interface{}{
						"error": errorMeta["error"],
					}
					c.JSON(http.StatusUnauthorized, errorJsonResponse)
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				}

			case http.StatusMethodNotAllowed:
				errorMeta, ok := ginErr.Meta.(map[string]interface{})
				if ok {
					errorJsonResponse := map[string]interface{}{
						"error": errorMeta["error"],
					}
					c.JSON(http.StatusMethodNotAllowed, errorJsonResponse)
				} else {
					c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
				}
			}
		}
	}
}