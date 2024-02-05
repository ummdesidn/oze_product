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
func (p Depp) GetDep(c *gin.Context) (results []Depp) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Dep")
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
func (p Employee) GetDepByID(c *gin.Context, idd string) (results []Employee) {
	id, _ := primitive.ObjectIDFromHex(idd) // แปลงค่า id
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Staff")
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
func Dep_New(c *gin.Context) {

	DepName := c.PostForm("DepName")

	now := time.Now()
	setTime := now.Unix()
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Dep")
	EmpNew := bson.M{"DepName": DepName, "Status": "1", "Checkpoint": setTime}
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
func Dep_Up(c *gin.Context) {
	PKkey := c.PostForm("PKKey")
	Old := c.PostForm("DepOld") // ชื่อเก่า
	Dep := c.PostForm("DepName")
	Status := c.PostForm("Status")
	fmt.Println("func Dep_UP : Old => ", Old)
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	now := time.Now()
	setTime := now.Unix()
	coll := client.Database(dbName).Collection("tbl_Dep")
	//ok => filter := bson.D{{"name", "Cup1"}}
	id, _ := primitive.ObjectIDFromHex(PKkey) // แปลงค่า id
	filter := bson.D{{"_id", id}}

	if Status == "0" {
		update := bson.D{{"$set", bson.D{{"Status", "0"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update Disable : ", result.MatchedCount)
	} else {
		update := bson.D{{"$set", bson.D{{"DepName", Dep}, {"Status", "1"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		// up tbl_Staff

		coll_staff := client.Database(dbName).Collection("tbl_Dep")
		filterStaff := bson.D{{"Dep", Old}}
		updateStaff := bson.D{{"$set", bson.D{{"Dep", Dep}}}}
		resultStaff, err := coll_staff.UpdateMany(ctx, filterStaff, updateStaff)
		if err != nil {
			panic(err)
		}
		fmt.Println("Update OK : ", result.MatchedCount, " || ", resultStaff)

	}
	cc := gin.H{"msg": "บันทึกข้อมูลเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func
