package handler

import (
	"crypto/tls"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

func Gettel(c *gin.Context) {

	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/gettel")
	if err != nil {

		c.String(http.StatusBadRequest, "err1")
	}
	tel := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, tel.String())
}

func Changetel(c *gin.Context) {
	tel := c.Param("tel")
	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/changetel/" + tel)
	if err != nil {

		c.String(http.StatusBadRequest, "err2")
	}
	telphote := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, telphote.String())

}
