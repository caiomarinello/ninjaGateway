package utils

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForwardRequest(c *gin.Context, url string, customHeaders map[string]string) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body: " + err.Error()})
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body)) // Reset the original body for middleware usage

	// Create a new HTTP request
	req, err := http.NewRequest(c.Request.Method, url, bytes.NewReader(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request: " + err.Error()})
		return
	}

	// Copy headers from the original request
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Add custom headers
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	// Forward the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to forward request: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	// Copy the response
	c.Status(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}
	io.Copy(c.Writer, resp.Body)
}