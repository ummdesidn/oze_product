package product

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

// / ดึงข้อมูล Dep
func (p AllHashtag) GetHashtag(c *gin.Context) (results []AllHashtag) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_hashtag")
	filter := bson.M{"Status": "ใช้งาน"}
	cursor, err := coll.Find(ctx, filter)

	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	/// append
	for _, result := range results {

		cursor.Decode(&result) // ถอดค่าก่อนส่งเข้า template

		//TXT_ID := result.PKKey.Hex()
	} // end for
	return results

} // end func
