package purchase

import (
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func URL_ShowSupplierAll(c *gin.Context) {
	fmt.Println("func : URL_ShowSupplierAll")
	alert := CheckAlert(c) ///ตรวจสอบอย่างเดียวไม่ต้องตั้งคค่า
	//staff := Employee{}
	//GetStaff := staff.GetStaff(c) // Staff
	bnk := Bank{}
	bank := bnk.GetBank()
	pv := Province{}
	getProvince := pv.GetProvince()
	pt := Supplyier{}
	getSupplier := pt.GetSupplier(c) // supplier.go
	cate := CateProduct{}
	CateProduct := cate.GetCateProduct(c)
	fmt.Println(getSupplier)
	cc := gin.H{"title": "หน้าแสดงผู้ผลิตสินค้า", "alert": alert, "getSupplier": getSupplier, "bank": bank, "getProvince": getProvince, "CateProduct": CateProduct}
	c.HTML(http.StatusOK, "purchase/supplier/home.tmpl.html", cc)

} // end func

func URL_SupplierDetail(c *gin.Context) {
	idd := c.PostForm("idd")

	alert := ""
	bnk := Bank{}
	bank := bnk.GetBank()
	pv := Province{}
	getProvince := pv.GetProvince()
	sup := Supplyier{}
	SupplierDetail := sup.GetSupplierByID(c, idd)
	cate := CateProduct{}
	CateProduct := cate.GetCateProductHaveID(c, idd)
	fmt.Println("ddd", SupplierDetail)
	cc := gin.H{"title": "หน้าแสดงรายละเอียดผู้ผลิตสินค้า", "alert": alert, "SupplierDetail": SupplierDetail, "bank": bank, "getProvince": getProvince, "CateProduct": CateProduct, "pkkey": idd} //
	c.HTML(http.StatusOK, "purchase/supplier/Detail.tmpl.html", cc)

} // end func
/*
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
