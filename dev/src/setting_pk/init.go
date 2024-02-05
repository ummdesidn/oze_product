package setting

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
	// Staff Start
	case "staff":
		URL_Staff(c) // url.go
	case "showStaff":
		URL_StaffDetail(c) // url.go
		// Staff End
		// Dep Start
	case "showDep":
		URL_Dep(c)
		// Del End
	case "datatable":
		URL_DataTable(c) // url.go

		/// กลุ่มตั้งค่าสินค้า
	case "ProductCate":
		URL_ProductCate(c) // url.go
		/// กลุ่มตั้งค่า hashtag
	case "ProductHashtag":
		URL_ProductHashtag(c) // url.go
		//// ราขละเอียด ธนาคาร
	case "Bank":
		URL_Bank(c) // url.go
		//// หน่วยนีบสินค้า
	case "ProductUnit":
		URL_ProductUnit(c)
	default:
		location := url.URL{Path: "/home"}
		c.Redirect(http.StatusFound, location.RequestURI())

	}

} // end func
