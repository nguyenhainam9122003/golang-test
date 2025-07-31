package utils

import (
	"github.com/gin-gonic/gin"
)

type PaginatedData struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Items interface{} `json:"items"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Respond(c *gin.Context, statusCode int, success bool, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	})
}

func HandleError(c *gin.Context, err error, statusCode int, message string) bool {
	if err != nil {
		if message == "" {
			message = err.Error()
		}
		Respond(c, statusCode, false, message, nil)
		return true
	}
	return false
}
