package formNow

import (
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func URL_Input(c *gin.Context) {
	alert := ""
	cc := gin.H{"title": "inputForm", "alert": alert}
	c.HTML(http.StatusOK, "form/input/home.tmpl.html", cc)

} // end func
func URL_DataTable(c *gin.Context) {
	alert := ""
	cc := gin.H{"title": "DataTable", "alert": alert}
	c.HTML(http.StatusOK, "form/datatable/home.tmpl.html", cc)

} // end func
