package purchase

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CountColumn ใช้นับจำนวน Column
func AjaxNow(c *gin.Context) {
	fmt.Println("Func : AjaxNow ")
	idd := c.PostForm("menu")
	switch idd {
	case "goDistrict":
		GetDistrict(c) //province.go
	case "goTambon":
		GetTambon(c) //province.go
	case "goPostcode":
		GetPostcode(c) //province.go
	case "NewSupplier":
		NewSupplier(c) //supplier.go
	case "UpSupplier":
		UpSupplier(c) // supplier.go
	} // end func
}
