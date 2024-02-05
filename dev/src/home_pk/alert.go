package home

import (
	"github.com/gin-gonic/gin"
)

var AlertNow string

/// ตรวจสอบ Alert
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
