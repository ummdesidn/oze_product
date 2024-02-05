package product

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var AlertNow string

// / ตรวจสอบ Alert
func CheckAlert(c *gin.Context) string {
	var msg string
	if AlertNow == "" {
		msg = ""
	} else {
		msg = AlertNow
		AlertNow = ""
	}
	return msg
}

/*
script go to page
location := url.URL{Path: "/purchase/3050"}
c.Redirect(http.StatusFound, location.RequestURI())
*/
// เมนูสำหรับแสดงข้อความและไปหน้าต่อไป
func GoNextPage(msg, nextpage string, c *gin.Context) {
	if msg == "" {
		AlertNow = ""
	} else {
		AlertNow = msg
	}
	location := url.URL{Path: nextpage}
	c.Redirect(http.StatusFound, location.RequestURI())
}
