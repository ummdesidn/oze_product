package purchase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// / ดึงข้อมูลจังหวัด
func (p Province) GetProvince() (results []Province) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Province")
	filter := bson.M{} //ProvinceThai
	//cursor, err := coll.Find(ctx, filter)
	cursor, err := coll.Distinct(ctx, "ProvinceThai", filter)

	if err != nil {
		panic(err)
	}
	var tmp Province
	for i, result := range cursor {

		fmt.Println(i, " -> ", result)   //แสดงข้อมูลถูกต้อง
		str := fmt.Sprintf("%v", result) ///
		tmp.ProvinceThai = str
		results = append(results, tmp)

	} // end for
	////

	return results
} // end func
// / ดึงอำเภอ
func GetDistrict(c *gin.Context) {

	var val_DistrictThai string
	var a_DistrictThai []string

	tmpProvince := c.PostForm("val_sel")
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Province")
	filter := bson.M{"ProvinceThai": tmpProvince}
	//cursor, err := coll.Find(ctx, filter)
	cursor, err := coll.Distinct(ctx, "DistrictThai", filter)

	//defer cursor.Close(ctx)
	if err != nil {
		panic(err)
	}
	for _, result := range cursor {

		str := fmt.Sprintf("%v", result) ///
		val_DistrictThai = str
		a_DistrictThai = append(a_DistrictThai, val_DistrictThai)
		//	fmt.Println(i, " -> ", result)   //แสดงข้อมูลถูกต้อง
		//	results = append(results, tmp)
		//	fmt.Println("tambon : ", a_DistrictThai)
	} // end for

	cc := gin.H{"Status": "ok", "retNow": a_DistrictThai}
	c.JSON(200, cc)
} // end func
// / ดึงตำบล
func GetTambon(c *gin.Context) {

	var val_Tambon string
	var a_Tambon []string

	tmpTambon := c.PostForm("val_sel")
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Province")
	filter := bson.M{"DistrictThai": tmpTambon}
	//cursor, err := coll.Find(ctx, filter)
	cursor, err := coll.Distinct(ctx, "TambonThaiShort", filter)

	if err != nil {
		panic(err)
	}
	for _, result := range cursor {

		str := fmt.Sprintf("%v", result) ///
		val_Tambon = str
		a_Tambon = append(a_Tambon, val_Tambon)
		//	fmt.Println(i, " -> ", result)   //แสดงข้อมูลถูกต้อง
		//	results = append(results, tmp)
		//	fmt.Println("tambon : ", a_DistrictThai)
	} // end for

	cc := gin.H{"Status": "ok", "retNow": a_Tambon}
	c.JSON(200, cc)
} // end func
// / ดึงเลขไปรษณีย์
func GetPostcode(c *gin.Context) {

	var val_Postcode string
	//var a_Postcode []string

	tmpPostcode := c.PostForm("val_sel")
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Province")
	filter := bson.M{"TambonThaiShort": tmpPostcode}
	//cursor, err := coll.Find(ctx, filter)
	cursor, err := coll.Distinct(ctx, "PostCodeMain", filter)

	if err != nil {
		panic(err)
	}
	for _, result := range cursor {

		str := fmt.Sprintf("%v", result) ///
		val_Postcode = str
		//a_Postcode = append(a_Postcode, val_Postcode)
		//	fmt.Println(i, " -> ", result)   //แสดงข้อมูลถูกต้อง
		//	results = append(results, tmp)
		//	fmt.Println("tambon : ", a_DistrictThai)
	} // end for

	cc := gin.H{"Status": "ok", "retNow": val_Postcode}
	c.JSON(200, cc)
} // end func
//////////////////////
/////////////////////
////////////////////////////////////
/////// ด้านล่างขอเดิม
/*
/// getAllEmployee

func (p Dep) getDep() (results []Dep) {

	ctx, client, dbName := dbConn()

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

	} // end for

	fmt.Println("ch 1 ", results)
	return results

}
*/
//// tmp Code
/*
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"registertime", bson.D{{"$gt", -1}}}},
				bson.D{{"registertime", bson.D{{"$lt", 1}}}},
			}},
	} // end filter
*/

//filter := bson.D{{"name", bson.D{{"$in", "^1"}}}}
//filter := bson.M{"level": "1"}

//text
/*
	model := mongo.IndexModel{Keys: bson.D{{"name", "text"}}}
	name, err := coll.Indexes().CreateOne(ctx, model)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name of index created: " + name)
	filter := bson.D{{"$text", bson.D{{"$search", "1"}}}}
*/
// end text

//filter := bson.M{"name": "emp001"}
