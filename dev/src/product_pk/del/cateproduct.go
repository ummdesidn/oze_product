package setting

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// / ดึงข้อมูล ทั้งหมด
func (p CateProduct) GetCateProduct(c *gin.Context) (results []CateProduct) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_ProductCate")
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
func (p CateProduct) GetCateProductByID(c *gin.Context, idd string) (results []CateProduct) {
	id, _ := primitive.ObjectIDFromHex(idd) // แปลงค่า id
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_ProductCate")
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

// เพิ่มข้อมูล หมวดสินค้า
func ProductCate_New(c *gin.Context) {
	//now := time.Now()
	//setTime := now.Unix()

	CateProduct := c.PostForm("CateProduct")
	Status := c.PostForm("Status")

	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_ProductCate")
	EmpNew := bson.M{"CateName": CateProduct, "CateStatus": Status}
	result, err := coll.InsertOne(ctx, EmpNew) // {"Contact}
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert OK : ", result)

	///// ส่งสถานะกลับ
	cc := gin.H{"msg": "บันทึกข้อมูลเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func

// // ปรับปรุงข้อมูลตาม PKkry
func CateProduct_Up(c *gin.Context) {
	PKkey := c.PostForm("PKKey1")
	CateName := c.PostForm("CateName")
	Status := c.PostForm("Status")
	fmt.Println("PKkey :", PKkey, " CateName : ", CateName, " Status : ", Status)
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_ProductCate")
	//ok => filter := bson.D{{"name", "Cup1"}}
	id, _ := primitive.ObjectIDFromHex(PKkey) // แปลงค่า id
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"CateName", CateName}, {"Status", Status}}}}
	result, _ := coll.UpdateOne(ctx, filter, update)
	/*
		// up tbl_Product
		coll_staff := client.Database(dbName).Collection("tbl_Staff")
		filterStaff := bson.D{{"Dep", Old}}
		updateStaff := bson.D{{"$set", bson.D{{"Dep", Dep}}}}
		resultStaff, err := coll_staff.UpdateMany(ctx, filterStaff, updateStaff)
		if err != nil {
			panic(err)
		}

	*/
	fmt.Println("Update OK : ", result.MatchedCount, " || ")
	cc := gin.H{"msg": "บันทึกข้อมูลเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func
