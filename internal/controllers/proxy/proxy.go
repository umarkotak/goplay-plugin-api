package cProxy

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/goplay-plugin-api/internal/lib/utils"
)

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func GetProxy(c *gin.Context) {
	targetURI := strings.ReplaceAll(c.Request.RequestURI, "/proxy/", "")
	logrus.Info(targetURI)

	client := &http.Client{}
	req, err := http.NewRequest("GET", targetURI, nil)

	resp, err := client.Do(req)
	if err != nil {
	}

	defer resp.Body.Close()

	c.Writer.Header().Set("Content-Type", "application/json")
	io.Copy(c.Writer, resp.Body)

	utils.RenderResponse(c, resp.StatusCode, nil)
}
