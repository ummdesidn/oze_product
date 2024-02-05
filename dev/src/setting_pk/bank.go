package setting

import (
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// / ดึงข้อมูล Dep
func (p Bank) GetBank(c *gin.Context) (results []Bank) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Bank")
	filter := bson.M{}
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
// / ดึงข้อมูลเจ้าหน้าที่ตามค่า _id ยังไม่ใช้
func (p Bank) BankByID(c *gin.Context, idd string) (results []Bank) {
	id, _ := primitive.ObjectIDFromHex(idd) // แปลงค่า id
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Bank")
	filter := bson.M{"_id": id}
	cursor, err := coll.Find(ctx, filter)

	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	/// append
	for _, result := range results {

		cursor.Decode(&result) // ถอดค่าก่อนส่งเข้า template

	} // end for
	return results

} // end func
// เพิ่มข้อมูล ยังไม่ได้ทำ
func Bank_New(c *gin.Context) {

	Namee := c.PostForm("Namee")

	now := time.Now()
	setTime := now.Unix()
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Bank")
	EmpNew := bson.M{"BankName": Namee, "Status": "ใช้งาน", "Checkpoint": setTime}
	result, err := coll.InsertOne(ctx, EmpNew) // {"Contact}
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert OK : ", result)

	///// ส่งสถานะกลับ
	cc := gin.H{"msg": "เพิ่มข้อมูลเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func

// // ปรับปรุงข้อมูลตาม PKkry
func Bank_Up(c *gin.Context) {
	PKkey := c.PostForm("PKKey1")
	Old := c.PostForm("OldVal") // ชื่อเก่า
	Namee := c.PostForm("Namee")
	Status := c.PostForm("Status")
	fmt.Println("Up Detail: Old => ", Old, " new Val = ", Namee)
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	now := time.Now()
	setTime := now.Unix()
	coll := client.Database(dbName).Collection("tbl_Bank")
	//ok => filter := bson.D{{"name", "Cup1"}}
	id, _ := primitive.ObjectIDFromHex(PKkey) // แปลงค่า id
	filter := bson.D{{"_id", id}}

	if Status == "ระงับการใช้งาน" {
		update := bson.D{{"$set", bson.D{{"Status", "ระงับการใช้งาน"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update Disable : ", result.MatchedCount)
	} else {
		update := bson.D{{"$set", bson.D{{"BankName", Namee}, {"Status", "ใช้งาน"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		/*
			// up tbl_Staff

			coll_staff := client.Database(dbName).Collection("tbl_hashtag")
			filterStaff := bson.D{{"Dep", Old}}
			updateStaff := bson.D{{"$set", bson.D{{"HashtagName", Dep}}}}
			resultStaff, err := coll_staff.UpdateMany(ctx, filterStaff, updateStaff)
			if err != nil {
				panic(err)
			}*/
		fmt.Println("Update OK : ", result.MatchedCount)

	}
	cc := gin.H{"msg": "บันทึกข้อมูลเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func
