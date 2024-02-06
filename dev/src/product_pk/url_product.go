package product

import (
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func URL_ShowMainProduct(c *gin.Context) {
	alert := CheckAlert(c)

	prod := AllProduct{}
	GetProduct := prod.GetNewProductLast10(c)
	///fmt.Println("GetProduct :::: ", GetProduct)
	cc := gin.H{"title": "หน้าหลักแสดงสินค้า", "alert": alert, "GetProduct": GetProduct}
	c.HTML(http.StatusOK, "product/product/2-1_home.tmpl.html", cc)

} // end func

func URL_ShowProductAll(c *gin.Context) {
	alert := CheckAlert(c)
	//staff := Employee{}
	//GetStaff := staff.GetStaff(c) // Staff
	//okprod := AllProduct{}
	prod := TotalProduct{}
	GetProduct := prod.GetProduct(c) // product.go

	uni := AllUnitProduct{} // หน่วยนับสินค้า
	GetAllUnitProduct := uni.GetAllUnitProduct(c)
	CateP := CateProduct{}                           // ประเภทสินค้า
	CateProduct := CateP.GetCateProductWithEnable(c) // cateproduct.go

	hasht := AllHashtag{}
	Hashtag := hasht.GetHashtag(c) // hashtag.go
	//fmt.Println("CateProduct :: ", CateProduct)
	cc := gin.H{"title": "หน้าแสดงสินค้า", "alert": alert, "CateProduct": CateProduct, "GetAllUnitProduct": GetAllUnitProduct, "GetProduct": GetProduct, "Hashtag": Hashtag}
	c.HTML(http.StatusOK, "product/product/home.tmpl.html", cc)

} // end func
func URL_ProductDetail(c *gin.Context) {
	//alert := ""
	//staff := Employee{}
	//GetStaff := staff.GetStaff(c) // Staff
	hasht := AllHashtag{}
	Hashtag := hasht.GetHashtag(c) // hashtag.go
	prod := AllProduct{}
	GetProduct := prod.GetProductWithPK(c)

	uni := AllUnitProduct{} // หน่วยนับสินค้า
	GetAllUnitProduct := uni.GetAllUnitProduct(c)
	CateP := CateProduct{}                           // ประเภทสินค้า
	CateProduct := CateP.GetCateProductWithEnable(c) // cateproduct.go
	imgID := GetProduct[1].DateIn

	fmt.Println("imgID :: ", imgID)
	cc := gin.H{"title": "หน้าแสดงรายละเอียดสินค้า", "PKKEY": c.PostForm("idd"), "CateProduct": CateProduct, "GetAllUnitProduct": GetAllUnitProduct, "GetProduct": GetProduct, "imgID": imgID, "Hashtag": Hashtag}
	c.HTML(http.StatusOK, "product/product/Detail.tmpl.html", cc)

} // end func
/*
func URL_StaffDetail(c *gin.Context) {
	idd := c.Param("idd")

	alert := ""
	staff := Employee{}
	StaffDetail := staff.GetStaffByID(c, idd) // Staff
	fmt.Println("ddd", StaffDetail)
	cc := gin.H{"title": "พนักงาน", "alert": alert, "StaffDetail": StaffDetail}
	c.HTML(http.StatusOK, "setting/staff/staffDetail.tmpl.html", cc)

} // end func
// // Dep
func URL_Dep(c *gin.Context) {
	depp := Depp{}
	GetDep := depp.GetDep(c)
	cc := gin.H{"title": "แผนก", "GetDep": GetDep}
	c.HTML(http.StatusOK, "setting/dep/home.tmpl.html", cc)
} // end func

// //////////////////////////////////////////////////////////////////////
// / กลุ่ม สินค้า หน้าหลัก
func URL_ProductCate(c *gin.Context) {

	cate := CateProduct{}
	Cate := cate.GetCateProduct(c)
	cc := gin.H{"title": "หมวดสินค้า", "Cate": Cate}
	c.HTML(http.StatusOK, "setting/productCate/home.tmpl.html", cc)
} // end func
// / กลุ่ม สินค้า -> รายละเอียดเพิ่มเิตท
func URL_CateProductDetail(c *gin.Context) {
	idd := c.Param("idd")
	cate := CateProduct{}
	CateNow := cate.GetCateProductByID(c, idd)

	cc := gin.H{"title": "หมวดสินค้า", "CateNow": CateNow}
	c.HTML(http.StatusOK, "setting/productCate/Detail.tmpl.html", cc)
} // end func

// ///////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// // ไมได้ใช้
func URL_DataTable(c *gin.Context) {
	alert := ""
	cc := gin.H{"title": "DataTable", "alert": alert}
	c.HTML(http.StatusOK, "form/datatable/home.tmpl.html", cc)

} // end func
*/
