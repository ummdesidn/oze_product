package purchase

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// / ดึงข้อมูล Dep
func (p Bank) GetBank() (results []Bank) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Bank")
	filter := bson.M{"Status": "ใช้งาน"}
	opts := options.Find().SetSort(bson.D{{"BankName", 1}})
	cursor, err := coll.Find(ctx, filter, opts)

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
// / ดึงข้อมูลเจ้าหน้าที่ตามค่า _id ยังไม่ใช้
