package setting

import (
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func URL_Staff(c *gin.Context) {
	alert := ""
	staff := Employee{}
	GetStaff := staff.GetStaff(c) // Staff

	cc := gin.H{"title": "พนักงาน", "alert": alert, "GetStaff": GetStaff}
	c.HTML(http.StatusOK, "setting/staff/home.tmpl.html", cc)

} // end func
func URL_StaffDetail(c *gin.Context) {
	idd := c.Param("idd")

	alert := ""
	staff := Employee{}
	StaffDetail := staff.GetStaffByID(c, idd) // Staff
	depp := Depp{}
	shdep := depp.GetDep(c)
	fmt.Println("ddd", StaffDetail)
	cc := gin.H{"title": "พนักงาน", "alert": alert, "StaffDetail": StaffDetail, "shdep": shdep}
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

// // Hashtag
func URL_ProductHashtag(c *gin.Context) {
	hash := Hashtag{}
	hashtag := hash.GetHashtag(c)
	cc := gin.H{"title": "หมวด Hashtag", "hashtag": hashtag}
	c.HTML(http.StatusOK, "setting/productHashtag/home.tmpl.html", cc)

} // end func
func URL_HashtagDetail(c *gin.Context) {
	idd := c.Param("idd")
	hash := Hashtag{}
	hashtag := hash.HashtagByID(c, idd)

	cc := gin.H{"title": "หมวด Hashtag", "hashtag": hashtag}
	c.HTML(http.StatusOK, "setting/productHashtag/Detail.tmpl.html", cc)
} // end func
///// bank

func URL_Bank(c *gin.Context) {
	poin := Bank{}
	bank := poin.GetBank(c)
	cc := gin.H{"title": "หมวด ข้อมูลธนาคาร", "bank": bank}
	c.HTML(http.StatusOK, "setting/bank/home.tmpl.html", cc)

} // end func
func URL_BankDetail(c *gin.Context) {
	idd := c.Param("idd")
	poin := Bank{}
	bank := poin.BankByID(c, idd)

	cc := gin.H{"title": "หมวด ข้อมูลธนาคาร", "bank": bank}
	c.HTML(http.StatusOK, "setting/bank/Detail.tmpl.html", cc)
} // end func
// ///////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// // ไมได้ใช้
func URL_DataTable(c *gin.Context) {
	alert := ""
	cc := gin.H{"title": "DataTable", "alert": alert}
	c.HTML(http.StatusOK, "form/datatable/home.tmpl.html", cc)

} // end func
