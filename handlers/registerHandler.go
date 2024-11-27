package handlers

import (
	"errors"
	"net/http"

	comp "github.com/caiomarinello/ninjaGateway/components"
	repos "github.com/caiomarinello/ninjaGateway/repositories"
	"github.com/caiomarinello/ninjaGateway/utils"
	"github.com/gin-gonic/gin"
)

func HandleRegisterUser(registrar repos.Registrar[comp.User]) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			c.AbortWithError(http.StatusMethodNotAllowed, errors.New("method different than POST")).SetMeta(map[string]interface{}{
				"error": "method not allowed",
			})
			return
		}

		var newUser comp.User
		if err := c.BindJSON(&newUser); err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "invalid request body",
			})
			return
		}

		ok, err := utils.ValidateStructFields(c, newUser)
		if !ok {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := registrar.Register(newUser); err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(err.Error())).SetMeta(map[string]interface{}{
				"error": "failed to register user",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
	}
}