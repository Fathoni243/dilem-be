package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ResponseSuccessJson(c *gin.Context, message string, data interface{}) {
	
	if message == "" {
		message = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"success": true,
		"data": data,
	})
}

func ResponseValidationErrorJson(c *gin.Context, message string, detail interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"success": false,
		"data": detail,
	})
}

func ResponseValidatorErrorJson(c *gin.Context, err error) {
	errorMessages := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		errorMessages = append(errorMessages, errorMessage)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"detail": errorMessages,
	})
}

func ResponseErrorJson(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func ResponseDetailErrorJson(c *gin.Context, message string, detail interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
	   "message": message,
	   "success": false,
	   "data": detail,
	})
}

func ResponseWhenFailOrError(c *gin.Context, code int, err error)  {
	c.JSON(code, gin.H{
		"success": false,
		"message": err.Error(),
	 })
}