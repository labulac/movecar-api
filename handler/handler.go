package handler

import (
	"crypto/sha1"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"

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

func Get58(c *gin.Context) {

	test := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := test.R().Get("http://home.labulac.top:40080/get58")
	if err != nil {

		c.String(http.StatusInternalServerError, "远端不可达")
	}
	tel := gjson.Get(resp.String(), "msg")

	c.String(http.StatusOK, tel.String())
}

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{"123456", timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateUrl(ctx *gin.Context) bool {
	timestamp := ctx.Query("timestamp")
	nonce := ctx.Query("nonce")
	signature := ctx.Query("signature")
	echostr := ctx.Query("echostr")
	signatureGen := makeSignature(timestamp, nonce)

	//log.Println(timestamp, nonce, signature)

	if signatureGen != signature {
		return false
	}
	_, _ = ctx.Writer.WriteString(echostr)

	return true
}

func Wechat(ctx *gin.Context) {

	if !validateUrl(ctx) {
		log.Println("Wechat Service: This http request is not from wechat platform")
		return
	}
	log.Println("validateUrl Ok")
}

func Sub(c *gin.Context) {
	resp1, err := http.Get("https://ghproxy.com/https://gist.githubusercontent.com/labulac/683e328ec813a45999b5893194189ac2/raw/sub.txt")
	if err != nil {
		c.String(http.StatusBadRequest, "Get sub file error: %v", err)
		return
	}
	defer resp1.Body.Close()
	content, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Read sub file error: %v", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	var cleanedLines []string

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" && !strings.HasPrefix(trimmedLine, " ") && !strings.HasPrefix(trimmedLine, "#") {
			cleanedLines = append(cleanedLines, trimmedLine)
		}
	}

	str := strings.Join(cleanedLines, "|")

	c.String(http.StatusOK, string(str))

}
