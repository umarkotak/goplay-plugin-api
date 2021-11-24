package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	cPing "github.com/umarkotak/goplay-plugin-api/internal/controllers/ping"
	cProxy "github.com/umarkotak/goplay-plugin-api/internal/controllers/proxy"
	"github.com/umarkotak/goplay-plugin-api/internal/lib/utils"
)

func Start() {
	r := InitRouter()

	port := 8000
	logrus.Infof("goplay plugin api is listening on port :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/ping", cPing.GetPing)
	r.GET("/proxy/*target_path", cProxy.GetProxy)

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			utils.RenderResponse(c, 200, gin.H{
				"success": true,
				"data":    nil,
				"error":   "",
			})
			return
		}
		c.Next()
	}
}
