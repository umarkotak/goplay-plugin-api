package utils

import (
	"github.com/gin-gonic/gin"
)

func RenderResponse(c *gin.Context, statusCode int, body interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")

	c.JSON(statusCode, body)
}
