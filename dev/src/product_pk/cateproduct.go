package product

/*
* function แสดงข้อมูล กลุ่มสินค้า
 */
import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

// / ดึงข้อมูล ทั้งหมด   (p Supplyier) GetSupplier(c *gin.Context) (results []Supplyier)
func (p CateProduct) GetCateProductWithEnable(c *gin.Context) (results []CateProduct) {

	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_ProductCate")
	filter := bson.M{"Status": "ใช้งาน"}
	cursor, err := coll.Find(ctx, filter)

	defer cursor.Close(ctx)

	if err != nil {
		panic(err)
	}

	var result CateProduct
	for cursor.Next(ctx) {
		//var result Supplyier
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", result)
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
	return results

	/*
		db := dbConn()
		prep, _ := db.Prepare("SELECT  COALESCE(cate.idCategory	,''),   COALESCE(cate.CategoryName,''), COALESCE(cate.CategoryStatus,'')  " +
			" FROM tbl_category  cate  " +
			" WHERE cate.CategoryStatus = 1 ")
		row, err1 := prep.Query()
		if err1 != nil {
			panic(err1)
		}

		for row.Next() {
			var tmp ALLCateProduct
			if err := row.Scan(&tmp.PKKey, &tmp.CateName, &tmp.Status); err != nil {
				panic(err)
			}
			if tmp.Status == "1" {
				tmp.TxtStatus = "ใช้งาน"
			} else {
				tmp.TxtStatus = "ระงับการใช้งาน"
			}
			RetNow = append(RetNow, tmp)

		}
		defer row.Close()
		defer db.Close()
		return RetNow
	*/

} // end func
