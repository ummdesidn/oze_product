package formNow

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

/*
func MainPage(c *gin.Context) {
	msg := CheckAlert(c) // check Alert
	cc := gin.H{"title": "หน้าหลัก", "alert": msg}
	c.HTML(http.StatusOK, "home/home/home.tmpl.html", cc)
} // end  func
*/
////////////////////////////////////
/// ด้านล่างของเดิม
func CheckNow(c *gin.Context) {
	idd := c.Param("idd")

	switch idd {
	case "formInput":
		URL_Input(c) // url.go
	case "datatable":
		URL_DataTable(c) // url.go

	default:
		location := url.URL{Path: "/home"}
		c.Redirect(http.StatusFound, location.RequestURI())

	}

} // end func
func PostNow(c *gin.Context) {
	/*	menu := c.PostForm("menu")

		switch menu {
		case "showinven":
			Fn_1011(c) // pc
		case "showinven_printer":
			Fn_1012(c)
		case "showinven_scanner":
			Fn_1013(c)
		case "showinven_notebook":
			Fn_1014(c)
		case "showinven_aio":
			Fn_1015(c)
		case "Upinven":
			UpdateInvenByID(c) // inventory.go
		case "Newinven":
			InsertNewInven(c) // inventory.go
		}
	*/

}
