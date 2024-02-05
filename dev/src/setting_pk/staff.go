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

/// ดึงข้อมูลเจ้าหน้าที่ที่ใช้งาน
func (p Employee) GetStaff(c *gin.Context) (results []Employee) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Staff")
	filter := bson.M{}
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
/// ดึงข้อมูลเจ้าหน้าที่ตามค่า _id
func (p Employee) GetStaffByID(c *gin.Context, idd string) (results []Employee) {
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
// เพิ่มข้อมูลเจ้าหน้าที่
func Staff_New(c *gin.Context) {

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

//// ปรับปรุงข้อมูลเจ้าหน้าที่ตาม PKkry
func Staff_Up(c *gin.Context) {
	now := time.Now()
	setTime := now.Unix()
	StaffName := c.PostForm("StaffName")
	Dep := c.PostForm("Dep")

	JobLevel := c.PostForm("Level")
	password := c.PostForm("pwd")
	// en pwd

	hasher := md5.New()
	hasher.Write([]byte(password))
	enPass := hex.EncodeToString(hasher.Sum(nil))
	username := c.PostForm("usr")

	Tel := c.PostForm("Tel")
	Mobile := c.PostForm("Mobile")
	Email := c.PostForm("Email")
	PKkey := c.PostForm("PKkey")
	Status := c.PostForm("Status")
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Staff")
	//ok => filter := bson.D{{"name", "Cup1"}}
	id, _ := primitive.ObjectIDFromHex(PKkey) // แปลงค่า id
	filter := bson.D{{"_id", id}}

	if Status == "0" {
		update := bson.D{{"$set", bson.D{{"Status", "0"}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update OK : ", result.MatchedCount)
	} else {
		update := bson.D{{"$set", bson.D{{"StaffName", StaffName}, {"Dep", Dep}, {"JobLevel", JobLevel}, {"username", username},
			{"CheckPoint", setTime}, {"Contact.Tel", Tel}, {"Contact.Mobile", Mobile}, {"Contact.Email", Email}, {"Status", "1"}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update OK : ", result.MatchedCount)
		if password != "" {
			update := bson.D{{"$set", bson.D{{"password", enPass}}}}
			result, _ := coll.UpdateOne(ctx, filter, update)
			fmt.Println("Update OK : ", result.MatchedCount)

		}
	}
	cc := gin.H{"msg": "บันทึกข้อมูลพนักงานเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func
