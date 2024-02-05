package product

/*
* function หน่วยนับสินค้า
 */
import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

// / ดึงข้อมูล ทั้งหมด
func (p AllUnitProduct) GetAllUnitProduct(c *gin.Context) (RetNow []AllUnitProduct) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Unit")
	filter := bson.M{"Status": "ใช้งาน"}
	cursor, err := coll.Find(ctx, filter)

	defer cursor.Close(ctx)

	if err != nil {
		panic(err)
	}

	var result AllUnitProduct
	var results []AllUnitProduct
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
	fmt.Println("product/GetAllUnitProduct/results : Unit --> ", results)
	return results
	/*
		db := dbConn()
		prep, _ := db.Prepare("SELECT  COALESCE(unt.idUnit	,''),   COALESCE(unt.UnitName,''), COALESCE(unt.UnitStatus,'')  " +
			" FROM tbl_unit  unt  " +
			" WHERE unt.UnitStatus = 1 ")
		row, err1 := prep.Query()
		if err1 != nil {
			panic(err1)
		}

		for row.Next() {
			var tmp AllUnitProduct
			if err := row.Scan(&tmp.PKKey, &tmp.UnitName, &tmp.Status); err != nil {
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
