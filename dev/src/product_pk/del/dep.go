package setting

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// ดึงข้อมูล Dep
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
/// ดึงข้อมูลเจ้าหน้าที่ตามค่า _id ยังไม่ใช้
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

	StaffName := c.PostForm("StaffName")
	Dep := c.PostForm("Dep")

	JobLevel := c.PostForm("Level")
	password := c.PostForm("pwd")
	username := c.PostForm("usr")

	Tel := c.PostForm("Tel")
	Mobile := c.PostForm("Mobile")
	Email := c.PostForm("Email")
	// en pwd
	hasher := md5.New()
	hasher.Write([]byte(password))
	enPass := hex.EncodeToString(hasher.Sum(nil))

	now := time.Now()
	setTime := now.Unix()
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Staff")
	EmpNew := bson.M{"StaffName": StaffName, "Dep": Dep, "JobLevel": JobLevel, "username": username, "password": enPass, "Status": "1", "CheckPoint": setTime,
		"Contact": bson.M{"Tel": Tel, "Mobile": Mobile, "Email": Email}}
	result, err := coll.InsertOne(ctx, EmpNew) // {"Contact}
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert OK : ", result)

	///// ส่งสถานะกลับ
	cc := gin.H{"msg": "บันทึกข้อมูลพนักงานเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func

//// ปรับปรุงข้อมูลตาม PKkry
func Dep_Up(c *gin.Context) {
	PKkey := c.PostForm("PKKey")
	Old := c.PostForm("DepOld") // ชื่อเก่า
	Dep := c.PostForm("DepName")
	Status := c.PostForm("Status")
	fmt.Println("func Dep_UP : Old => ", Old)
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Dep")
	//ok => filter := bson.D{{"name", "Cup1"}}
	id, _ := primitive.ObjectIDFromHex(PKkey) // แปลงค่า id
	filter := bson.D{{"_id", id}}

	if Status == "0" {
		update := bson.D{{"$set", bson.D{{"Status", "0"}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update Disable : ", result.MatchedCount)
	} else {
		update := bson.D{{"$set", bson.D{{"DepName", Dep}, {"Status", "1"}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		// up tbl_Staff

		coll_staff := client.Database(dbName).Collection("tbl_Staff")
		filterStaff := bson.D{{"Dep", Old}}
		updateStaff := bson.D{{"$set", bson.D{{"Dep", Dep}}}}
		resultStaff, err := coll_staff.UpdateMany(ctx, filterStaff, updateStaff)
		if err != nil {
			panic(err)
		}
		fmt.Println("Update OK : ", result.MatchedCount, " || ", resultStaff)

	}
	cc := gin.H{"msg": "บันทึกข้อมูลพนักงานเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func
