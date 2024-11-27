package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// If there are any errors, returns "Field errors: tag" in response body
func ValidateStructFields(c *gin.Context, structToValidate interface{}) (bool, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(structToValidate)
	if err != nil {
		fieldErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			fieldErrors[err.Field()] = err.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"Field errors": fieldErrors})
		return false, err
	}
	return true, nil
}
