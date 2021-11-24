package cProxy

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	req.Header.Add("cookie", "_gcl_au=1.1.1618906671.1637733228; session=IKDQWQUO2H2HCAQ23WJLRV52QDCATQ7GRMVN7NMBRFJCJSI2XCUSU2NMZRRIWKEI65G4V5EW5OHH5YI5KZZCDDMZPEMX3MHMT7DO6NA; _ga=GA1.3.1827307957.1637733228; _gid=GA1.3.804832792.1637733228; _gat_UA-162768158-1=1; _gorilla_csrf=MTYzNzczMzIyN3xJbEJoTUVzNFZUaHRVR3BqVDFWVmNqaFVURnBSTjBGd1FrVkJWekIwUjBwb0t6bEJPVmN2ZUZvck4zTTlJZ289fNGU9-Hjqv23COMlcIRRY6527nCuH3RKIeul6GSccsGQ; _fbp=fb.2.1637733228409.35417593; WZRK_G=640cc442501242c5b67b0b5c0407949a; afUserId=5152a5f0-20ec-47d4-9ec9-abbd77cba1ac-p; AF_SYNC=1637733229987; gp_daily_session=c390cacd-a9c6-452e-bcc3-0d7deb1714e7; WZRK_S_R5Z-568-WK5Z=%7B%%22p%%22%%3A3%2C%%22s%%22%3A1637733229%2C%%22t%%22%3A1637733237%7D; gp_fgp=5fef832d66288de65da6773ad50148c7%3Aaec34692b6c95c96eb46602d5e07a8f1a0acd0e7f7afe9c46206280e38cc21d6'")

	resp, err := client.Do(req)
	if err != nil {
	}

	defer resp.Body.Close()

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	io.Copy(c.Writer, resp.Body)
	return
}
