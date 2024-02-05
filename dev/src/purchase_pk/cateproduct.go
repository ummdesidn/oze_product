package purchase

import (
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

// / ดึงข้อมูล ทั้งหมด
func (p CateProduct) GetCateProduct(c *gin.Context) (results []CateProduct) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	coll := client.Database(dbName).Collection("tbl_ProductCate")

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

// / ดึงข้อมูล ทั้งหมด ด้วย idd
func (p CateProduct) GetCateProductHaveID(c *gin.Context, idd string) (results []CateProduct) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	/*
		coll := client.Database(dbName).Collection("tbl_Supplier")
		filter := bson.M{"_id": idd}
		cursor, _ := coll.Find(ctx, filter)
		defer cursor.Close(ctx)
	*/
	// ข้อมูล หมวสินค้า
	coll_Product := client.Database(dbName).Collection("tbl_ProductCate")
	filter_Product := bson.M{"Status": "ใช้งาน"}
	cursor_Product, err := coll_Product.Find(ctx, filter_Product)
	//var Cate []string
	defer cursor_Product.Close(ctx)
	if err = cursor_Product.All(ctx, &results); err != nil {
		panic(err)
	}
	/*
		/// append
		for _, result := range results {

			cursor_Product.Decode(&result) // ถอดค่าก่อนส่งเข้า template

			//TXT_ID := result.PKKey.Hex()
		} // end for
	*/
	/// loop
	var result CateProduct
	for cursor_Product.Next(ctx) {
		//var result Supplyier
		if err := cursor_Product.Decode(&result); err != nil {
			log.Fatal(err)
		}
		////// ทดสอบว่า ถ้าไม่มีค่าให้ append

		results = append(results, result)
	}
	if err := cursor_Product.Err(); err != nil {
		log.Fatal(err)
	}

	return results
	/*
		/// query
		lookupStage := bson.D{{"$lookup", bson.D{{"from", "tbl_Supplier"}, {"localField", "Dep"}, {"foreignField", "_id"}, {"as", "Dep_info"}}}}
		unwindStage := bson.D{{"$unwind", bson.D{{"path", "$Dep_info"}, {"preserveNullAndEmptyArrays", false}}}}
		// {"as", "Dep_info"} {"path", "$Dep_info"} คือการกไหนดค่าตัวแปรที่ใช้อ้างถึง
		showLoadedCursor, err := coll.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})
		if err != nil {
			panic(err)
		}
		var result CateProduct
		var tmpCount int = 1
		for showLoadedCursor.Next(ctx) {

			if err := showLoadedCursor.Decode(&result); err != nil {
				log.Fatal(err)
			}
			result.Test_Status = tmpCount
			tmpCount++
			fmt.Println(result.Test_Status)
			results = append(results, result)
		}
	*/
	/*
		if err = showLoadedCursor.All(ctx, &results); err != nil {
			panic(err)
		}
	*/
	//	fmt.Printf("Tea: %v \nToppings: %v \nPrice: $%v \n\n", result.Name, result.Dep1.DepName, result.Level)

	/*
		ของเดิมใช้งานได้
		coll := client.Database(dbName).Collection("tbl_ProductCate")
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


	*/
	//	return results

} // end func
/*
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

	fmt.Println("Update OK : ", result.MatchedCount, " || ")
	cc := gin.H{"msg": "บันทึกข้อมูลเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func
*/
