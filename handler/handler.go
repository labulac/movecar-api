package handler

import (
	"crypto/tls"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/ping")
	if err != nil {

		c.String(http.StatusInternalServerError, "远端不可达")
	}

	c.String(http.StatusOK, resp.String())
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

func Gettel(c *gin.Context) {

	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/gettel")
	if err != nil {

		c.String(http.StatusInternalServerError, "远端不可达")
	}
	tel := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, tel.String())
}

func Changetel(c *gin.Context) {
	tel := c.Param("tel")
	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/changetel/" + tel)
	if err != nil {

		c.String(http.StatusInternalServerError, "远端不可达")
	}
	telphote := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, telphote.String())

}

func Change58status(c *gin.Context) {
	status := c.Param("status")
	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/change58/" + status)
	if err != nil {

		c.String(http.StatusInternalServerError, "远端不可达")
	}
	telphote := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, telphote.String())

}

func CashSignin58(c *gin.Context) {

	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/58tc")
	if err != nil {

		c.String(http.StatusInternalServerError, "远端不可达")
	}
	telphote := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, telphote.String())

}

func Jimi(c *gin.Context) {

	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/jimi")
	if err != nil {

		c.String(http.StatusInternalServerError, "远端不可达")
	}
	telphote := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, telphote.String())

}
