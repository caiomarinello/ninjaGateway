package middleware

import (
	"log"
	"net/http"

	"github.com/caiomarinello/ninjaGateway/auth"
	"github.com/caiomarinello/ninjaGateway/components"
	"github.com/gin-gonic/gin"
)

// Middleware function to check if a user is logged in, adding the user data to the gin context if so.
func AuthenticateSession(store auth.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "sessionCookie")
		if err != nil {
			log.Println("error: no sessionCookie store found")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		sessionUser, ok := session.Values["user"].(components.User)
		if !ok || !sessionUser.Authenticated {
			log.Println("not authenticated")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
			return
		}

		// If authentication succeeds, user data is set to the context
		c.Set("sessionUser", sessionUser)
		c.Next()
	}
}

const (
	adminRole = "admin"
)

// Middleware function to check if a user is logged in and is an admin, adding the user data to the gin context if so.
func AuthenticateAdminSession(store auth.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "sessionCookie")
		if err != nil {
			log.Println("error: no sessionCookie store found")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		sessionUser, ok := session.Values["user"].(components.User)
		if !ok || !sessionUser.Authenticated {
			log.Println("not authenticated")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
			return
		}

		if sessionUser.Role != adminRole {
			log.Println("user does not have permission to access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no permission"})
			return
		}

		// If authentication succeeds, user data is set to the context
		c.Set("sessionUser", sessionUser)
		c.Next()
	}
}
