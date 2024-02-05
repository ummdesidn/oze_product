package product

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// //////////////////////////////////
// / ด้านล่างของเดิม
func CheckNow(c *gin.Context) {
	fmt.Println("product.go => CheckNow => ", c.Param("idd"))
	idd := c.Param("idd")

	switch idd {
	case "all":
		URL_ShowProductAll(c) // prodyct_pk/url.go
	case "productmain":
		URL_ShowMainProduct(c)
	// Staff Start

	default:
		//location := url.URL{Path: "/home"}
		//c.Redirect(http.StatusFound, location.RequestURI())

	}

} // end func
