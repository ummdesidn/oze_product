package product

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CountColumn ใช้นับจำนวน Column
func AjaxNow(c *gin.Context) {
	fmt.Println("Func : Product AjaxNow ")
	idd := c.PostForm("menu")
	switch idd {
	case "goSupplier":
		GetSupplierWithCate(c) //Supplier.go
	case "NewProduct":
		NewProduct(c) //product.go
	case "UpProduct":
		UpProduct(c) //product.go
	case "UpImg":
		UpImg(c) //product.go

	case "UpPriceProduct":
		UpPriceProduct(c) //product.go
	case "del_img":
		DelImgNow(c) //product.go

	} // end func
} // end func
