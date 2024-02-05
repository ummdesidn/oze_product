package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	msg := CheckAlert(c) // check Alert
	cc := gin.H{"title": "หน้าหลัก", "alert": msg}
	c.HTML(http.StatusOK, "home/home/home.tmpl.html", cc)
} // end  func
