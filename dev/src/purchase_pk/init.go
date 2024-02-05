package purchase

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// //////////////////////////////////
// / ด้านล่างของเดิม
func CheckNow(c *gin.Context) {
	idd := c.Param("idd")

	switch idd {
	case "supplier":
		URL_ShowSupplierAll(c) // purchase_pk/url.go
	// Staff Start

	default:
		location := url.URL{Path: "/home"}
		c.Redirect(http.StatusFound, location.RequestURI())

	}

} // end func
