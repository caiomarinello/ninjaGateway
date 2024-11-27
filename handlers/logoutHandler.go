package handlers

import (
	"errors"
	"net/http"

	"github.com/caiomarinello/ninjaGateway/auth"
	comp "github.com/caiomarinello/ninjaGateway/components"
	"github.com/caiomarinello/ninjaGateway/db"
	rep "github.com/caiomarinello/ninjaGateway/repositories"
	"github.com/gin-gonic/gin"
)

func HandleLogout(store auth.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "sessionCookie")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		sessionID := session.ID
		// Cookie probably expired
		if sessionID == "" {
			return
		}

		dbConn := db.OpenSqlConnection()
		defer dbConn.Close()
		err = rep.DeleteSession(dbConn, sessionID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to delete session during logout",
			})
			return
		}

		session.Values["user"] = comp.User{}
		session.Options.MaxAge = -1
		
		err = session.Save(c.Request, c.Writer)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to update session during logout",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	}
}
