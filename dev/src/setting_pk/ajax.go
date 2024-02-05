package setting

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// session ok
)

func AjaxNow(c *gin.Context) {
	fmt.Println("Func : AjaxNow ")
	idd := c.PostForm("menu")
	switch idd {
	case "staffnew":
		Staff_New(c) //setting_pk/staff.go
	case "staffup":
		Staff_Up(c) //setting_pk/staff.go
		// Dep Start
	case "DepUp":
		Dep_Up(c) // setting_pk/dep.go
	case "DepNew":
		Dep_New(c) // setting_pk/dep.go
		//Dep End

		///// หมวดสินค้า
	case "ProductCateNew":
		ProductCate_New(c) // setting_pk/cateproduct.go
	case "CateProductup":
		CateProduct_Up(c) // setting_pk/cateproduct.go

		/// หมวด Hashtag
	case "HashtagUp":
		Hashtag_Up(c)
	case "HashtagNew":
		Hashtag_New(c) // setting_pk/cateproduct.go

		///// หมวดธนาคาร
	case "BankUp":
		Bank_Up(c)
	case "BankNew":
		Bank_New(c) // setting_pk/bank.go
		//// หน่วยนับสินค้า
	case "ProductUnitNew":
		UnitProduct_New(c) // setting_pk/unitProduct.go
	case "ProductUnitUp":
		UnitProduct_Up(c) // setting_pk/unitProduct.go

	} // end func
}
