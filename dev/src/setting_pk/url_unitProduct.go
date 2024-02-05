package setting

import (
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// // Hashtag
func URL_ProductUnit(c *gin.Context) {
	hash := UnitProduct{}
	unitProduct := hash.GetUnitProduct(c)
	cc := gin.H{"title": "หมวด หน่วยนับสินค้า", "unitProduct": unitProduct}
	c.HTML(http.StatusOK, "setting/productUnit/home.tmpl.html", cc)

} // end func
func URL_ProductUnitDetail(c *gin.Context) {
	idd := c.Param("idd")
	hash := UnitProduct{}
	unitProduct := hash.UnitProductByID(c, idd)

	cc := gin.H{"title": "หมวด หน่วยนับสินค้า", "unitProduct": unitProduct}
	c.HTML(http.StatusOK, "setting/productUnit/Detail.tmpl.html", cc)
} // end func
