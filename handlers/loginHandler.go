package handlers

import (
	"errors"
	"net/http"

	"github.com/caiomarinello/ninjaGateway/auth"
	"github.com/caiomarinello/ninjaGateway/components"
	repos "github.com/caiomarinello/ninjaGateway/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(store auth.SessionStore, userFetcher repos.UserFetcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there isn't a session called sessionCookie, it will create one.
		session, err := store.Get(c.Request, "sessionCookie")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var loginBodyInf components.User
		if err := c.BindJSON(&loginBodyInf); err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "invalid request body",
			})
			return
		}
		// User validation
		userFromDb, err := userFetcher.FetchUserByEmail(loginBodyInf.Email)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "email not found",
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(loginBodyInf.Password))
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "invalid password",
			})
			return
		}
		user := &components.User{
			Id:				userFromDb.Id,
			Role:          	userFromDb.Role,
			Authenticated: 	true,
		}

		session.Values["user"] = user
		err = session.Save(c.Request, c.Writer)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to save session",
			})
			return
		}
		c.JSON(http.StatusFound, gin.H{"sucess": "Login sucesseful"})
		c.Redirect(http.StatusFound, "/index")
	}
}