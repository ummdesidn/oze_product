package product

/*
* function แสดงข้อมูล Supplier
 */
import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

// / ดึงข้อมูลหน้ารวม
func GetSupplierWithCate(c *gin.Context) {

	//id, _ := primitive.ObjectIDFromHex(idd) // แปลงค่า id
	typ := c.PostForm("val_sel")
	fmt.Println("typ :: ", typ)
	//	typ = "เมล็ดพันธ์"
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Supplier") // tbl_ProductCate  tbl_Supplier
	filter := bson.M{"ProductCate": typ}
	cursor, _ := coll.Find(ctx, filter)

	defer cursor.Close(ctx)

	var val_SupName string
	var A_SupName []string
	var val_id string
	var A_id []string

	var result AllSupplier // CateProduct AllSupplier
	//var results []AllSupplier // AllSupplier

	for cursor.Next(ctx) {
		//var result Supplyier
		if err := cursor.Decode(&result); err != nil {
			panic(err)
			log.Fatal(err)
		}

		//result.Address = "ssssss"
		val_SupName = result.SupName
		A_SupName = append(A_SupName, val_SupName)
		// id
		val_id = result.PKKey.String()
		A_id = append(A_id, val_id)

		//results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	///// ค้นหาข้อมูลสินค้าทดแทน ที่มีหมวดเดเียวกัน
	coll_productreplace := client.Database(dbName).Collection("tbl_Product") // tbl_ProductCate  tbl_Supplier
	filter_productreplace := bson.M{"ProductCate": typ, "Status": "ใช้งาน"}  //
	cursor_productreplace, _ := coll_productreplace.Find(ctx, filter_productreplace)
	var val_productreplace string
	var A_productreplace []string

	var result_productreplace AllProduct // CateProduct AllSupplier
	for cursor_productreplace.Next(ctx) {
		//var result Supplyier
		if err := cursor_productreplace.Decode(&result_productreplace); err != nil {
			panic(err)
			log.Fatal(err)
		}

		//result.Address = "ssssss"
		val_productreplace = result_productreplace.ProductName
		A_productreplace = append(A_productreplace, val_productreplace)

	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	defer cursor_productreplace.Close(ctx)

	/////// return

	fmt.Println("val_supname : ", A_SupName, " || A_productreplace : ", A_productreplace)

	c.JSON(200, gin.H{"sup_name": A_SupName, "PKKey": A_id, "productreplace": A_productreplace})

	/*
		///// ด้านล่างของ MySql
		idd := c.PostForm("val_sel")
		db := dbConn()
		prep, _ := db.Prepare("SELECT  DISTINCT COALESCE(sup.idSupplier,''),COALESCE(sup.SupplierID,''),   COALESCE(sup.SupplierName,'')" +
			" FROM tbl_supplier sup , tbl_suppliercategory supCate " +
			" WHERE sup.idSupplier = supCate.idSupplier AND sup.Status = 1 AND supCate.idCategory = ?  ")

		row, err1 := prep.Query(idd)
		if err1 != nil {
			panic(err1)
		}

		for row.Next() {
			if err := row.Scan(&val_pk, &val_supid, &val_supname); err != nil {
				panic(err)
			}
			a_pk = append(a_pk, val_pk)
			a_supid = append(a_supid, val_supid)
			a_supname = append(a_supname, val_supname)

		}
		defer row.Close()
		defer db.Close()
		fmt.Println("val_supname : ", a_supname)
		c.JSON(200, gin.H{"pk": a_pk, "sid": a_supid, "name": a_supname})
	*/
} // ebd func
// / ดึงข้อมูล Supplier ที่หน้ารายละเอียดสินค้า โดยจะแสดงข้อมูลตามประเภทของกลุ่มสินค้า
/*func (pt *AllSupplier) GetSupplierWithCateInDetailProduct(c *gin.Context, pk string) (RET []AllSupplier) {

		fmt.Println(" func: GetSupplierWithCateInDetailProduct :: ", pk)
		db := dbConn()
		prep, _ := db.Prepare("SELECT  DISTINCT COALESCE(sup.idSupplier,''),   COALESCE(sup.SupplierName,'') " +
			" FROM tbl_supplier sup , tbl_suppliercategory supCate " +
			" WHERE sup.idSupplier = supCate.idSupplier AND sup.Status = 1 AND supCate.idCategory = ?  ")

		row, err1 := prep.Query(pk)
		if err1 != nil {
			panic(err1)
		}

		//	var val_pk string
		//	var val_supid string
		//	var val_supname string

		//	var a_pk []string
		//	var a_supid []string
		//	var a_supname []string

		var tmp AllSupplier
		for row.Next() {
			if err := row.Scan(&tmp.PKKey, &tmp.SupplyierName); err != nil {
				panic(err)
			}
			//	a_pk = append(a_pk, val_pk)
			//	a_supid = append(a_supid, val_supid)
			//	a_supname = append(a_supname, val_supname)
			RET = append(RET, tmp)

		}
		defer row.Close()
		defer db.Close()
		return RET
		//fmt.Println("val_supname : ", a_supname)
		//c.JSON(200, gin.H{"pk": a_pk, "sid": a_supid, "name": a_supname})

} // ebd func
*/
