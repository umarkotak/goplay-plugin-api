package cPing

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/goplay-plugin-api/internal/lib/utils"
)

func GetPing(c *gin.Context) {

	pingResponse := map[string]string{
		"ping": "pong",
	}
	utils.RenderResponse(c, 200, pingResponse)
}
