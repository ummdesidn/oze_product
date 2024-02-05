package purchase

import (
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewSupplier(c *gin.Context) {
	// 2566-7-16 ใช้งานได้ปกติ
	now := time.Now()
	setTime := now.Unix()
	/// 1 ข้อมูลจาก หน้าเว็บ+
	suptype := c.PostForm("suptype")
	supid := c.PostForm("supid")
	suptitlename := c.PostForm("titlename")

	supname := c.PostForm("supname")
	suptax := c.PostForm("suptax")
	supaddress := c.PostForm("supaddress")
	province := c.PostForm("province")
	amp := c.PostForm("amp")
	tam := c.PostForm("tam")
	postid := c.PostForm("postid")
	suptel := c.PostForm("suptel")
	supcontactname := c.PostForm("supcontactname")
	supcontacttel := c.PostForm("supcontacttel")
	bankbrand := c.PostForm("bankbrand")
	bankno := c.PostForm("bankno")
	bankaccount := c.PostForm("bankaccount")
	credit := c.PostForm("credit")
	cash := c.PostForm("cash")
	supcates := c.PostFormArray("supcates")
	suppro := c.PostForm("suppro")
	supcomment := c.PostForm("supcomment")

	var ar2 []string
	for _, s := range supcates {
		ar2 = append(ar2, s)

	}
	//// ตรวจสอบจำนวนที่มีอยู่ในระบบ
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Supplier")
	filter := bson.M{}
	//cursor, err := coll.Find(ctx, filter)
	count, _ := coll.CountDocuments(ctx, filter)
	count = count + 10

	//txt_runNo := fmt.Sprintf("%04d\n", int(count)) // ok

	//fmt.Println("count ::: ", count)

	// ok
	newSupplier := bson.M{"Status": "ใช้งาน", "Address": supaddress, "SupID": supid, "SupType": suptype, "AmpurID": amp, "PostID": postid, "ProvinceID": province, "SupContactName": supcontactname, "SupContactTel": supcontacttel,
		"Suptitlename": suptitlename, "SupName": supname, "SupTaxID": suptax, "TambonID": tam, "Comment": supcomment, "CheckPoint": setTime, "SupPromotion": suppro, "GenSID": count,
		"DateIn": setTime, "SupTel": suptel, "Bank": bson.M{"BankName": bankbrand, "BankNo": bankno, "BookBank": bankaccount}, "Credit": bson.M{"Date": credit, "Cash": cash}, "ProductCate": ar2}
	// connect DB
	//ctx, client, dbName := DBConn()
	//defer client.Disconnect(ctx)

	//coll := client.Database(dbName).Collection("tbl_Supplier")
	if _, err := coll.InsertOne(ctx, newSupplier); err != nil {
		panic(err)

	}
	fmt.Println("Insert OK  ")
	/// 3 reload
	GoNextPage("บันทึกข้อมูล Supplyier เรียบร้อยแล้ว", "/purchase/supplier", c)

} // end func

func (p Supplyier) GetSupplier(c *gin.Context) (results []Supplyier) {
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Supplier")
	filter := bson.M{}
	cursor, err := coll.Find(ctx, filter)

	defer cursor.Close(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println("Test For ")
	var tmpCount int = 1
	var result Supplyier
	for cursor.Next(ctx) {
		//var result Supplyier
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		result.Count = tmpCount
		tmpCount++
		fmt.Println(result.Count)
		//fmt.Printf("%+v\n", result)
		// แปลงวันลงทะเบียนจาก timestamp to text
		var tmpTxt_DateIn string
		tmpTxt_DateIn = fmt.Sprint(result.DateIn)
		DateInTXT, err := strconv.Atoi(tmpTxt_DateIn)
		if err != nil {
			panic(err)
		}
		result.RegisterDateTXT = Timestamptotime(DateInTXT)
		fmt.Println("result.DateIn -> ", result.DateIn, " ---- tmpTxt_DateIn :: ->", tmpTxt_DateIn, " ---- DateInTXT -> ", DateInTXT, " ---- result.RegisterDateTXT  :::: ", result.RegisterDateTXT)
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return results

} // ebd func

// / ดึงข้อมูลเจ้าหน้าที่ตามค่า _id ยังไม่ใช้
func (p Supplyier) GetSupplierByID(c *gin.Context, idd string) (results []Supplyier) {
	id, _ := primitive.ObjectIDFromHex(idd) // แปลงค่า id
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Supplier")
	filter := bson.M{"_id": id}
	cursor, _ := coll.Find(ctx, filter)

	defer cursor.Close(ctx)
	var result Supplyier
	for cursor.Next(ctx) {
		//var result Supplyier
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return results

} // end func

func UpSupplier(c *gin.Context) {
	// 2566-7-16 ใช้งานได้ปกติ
	now := time.Now()
	setTime := now.Unix()
	/// 1 ข้อมูลจาก หน้าเว็บ+
	PKkey := c.PostForm("PKkey")

	suptype := c.PostForm("suptype")
	supid := c.PostForm("supid")
	suptitlename := c.PostForm("titlename")
	supname := c.PostForm("supname")
	suptax := c.PostForm("suptax")
	supaddress := c.PostForm("supaddress")
	province := c.PostForm("province")
	amp := c.PostForm("amp")
	tam := c.PostForm("tam")
	postid := c.PostForm("postid")
	suptel := c.PostForm("suptel")
	supcontactname := c.PostForm("supcontactname")
	supcontacttel := c.PostForm("supcontacttel")
	bankbrand := c.PostForm("bankbrand")
	bankno := c.PostForm("bankno")
	bankaccount := c.PostForm("bankaccount")
	credit := c.PostForm("credit")
	cash := c.PostForm("cash")
	supcates := c.PostFormArray("supcates")
	suppro := c.PostForm("suppro")
	supcomment := c.PostForm("supcomment")
	Status := c.PostForm("Status")

	var ar2 []string
	for _, s := range supcates { /// หมวดสินค้า
		ar2 = append(ar2, s)

	}
	fmt.Println("AR2 :=> ", ar2)
	/// เริ่มบันทึก
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Supplier")
	//ok => filter := bson.D{{"name", "Cup1"}}
	id, _ := primitive.ObjectIDFromHex(PKkey) // แปลงค่า id
	filter := bson.D{{"_id", id}}

	if Status == "ระงับการใช้งาน" {
		update := bson.D{{"$set", bson.D{{"Status", "ระงับการใช้งาน"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update Disable : ", result.MatchedCount)
	} else {
		//update := bson.D{{"$set", bson.D{{"HashtagName", Namee}, {"Status", "ใช้งาน"}, {"Checkpoint", setTime}}}}
		//result, _ := coll.UpdateOne(ctx, filter, update)

		update := bson.D{{"$set", bson.D{{"Status", "ใช้งาน"}, {"Address", supaddress}, {"SupID", supid}, {"SupType", suptype}, {"AmpurID", amp}, {"PostID", postid}, {"ProvinceID", province}, {"SupContactName", supcontactname}, {"SupContactTel", supcontacttel}, {"Suptitlename", suptitlename}, {
			"SupName", supname}, {"SupTaxID", suptax}, {"TambonID", tam}, {"Comment", supcomment}, {"CheckPoint", setTime}, {"SupPromotion", suppro}, {
			"DateIn", setTime}, {"SupTel", suptel}, {"Bank", bson.D{{"BankName", bankbrand}, {"BankNo", bankno}, {"BookBank", bankaccount}}}, {"Credit", bson.D{{"Date", credit}, {"Cash", cash}}}, {"ProductCate", ar2}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		/*
			newSupplier := bson.M{"Status": "ใช้งาน", "Address": supaddress, "SupID": supid, "SupType": suptype, "AmpurID": amp, "PostID": postid, "ProvinceID": province, "SupContactName": supcontactname, "SupContactTel": supcontacttel,
			"SupName": supname, "SupTaxID": suptax, "TambonID": tam, "Comment": supcomment, "CheckPoint": setTime, "SupPromotion": suppro,
			"DateIn": setTime, "SupTel": suptel, "Bank": bson.M{"BankName": bankbrand, "BankNo": bankno, "BookBank": bankaccount}, "Credit": bson.M{"Date": credit, "Cash": cash}, "ProductCate": ar2}
		*/

		// up tbl_Staff
		/*
			coll_staff := client.Database(dbName).Collection("tbl_Supplier")
			filterStaff := bson.D{{"Dep", Old}}
			updateStaff := bson.D{{"$set", bson.D{{"HashtagName", Dep}}}}
			resultStaff, err := coll_staff.UpdateMany(ctx, filterStaff, updateStaff)
			if err != nil {
				panic(err)
			}
		*/
		fmt.Println("Update OK : ", result.MatchedCount)

	}

	////
	fmt.Println("Insert OK  ")
	/// 3 reload
	GoNextPage("บันทึกข้อมูล Supplyier เรียบร้อยแล้ว", "/purchase/supplier", c)

} // end func

/*

// เพิ่มข้อมูล ยังไม่ได้ทำ
func Hashtag_New(c *gin.Context) {

	Namee := c.PostForm("Namee")

	now := time.Now()
	setTime := now.Unix()
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_hashtag")
	EmpNew := bson.M{"HashtagName": Namee, "Status": "ใช้งาน", "Checkpoint": setTime}
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
func Hashtag_Up(c *gin.Context) {
	PKkey := c.PostForm("PKKey1")
	Old := c.PostForm("DepOld") // ชื่อเก่า
	Namee := c.PostForm("Namee")
	Status := c.PostForm("Status")
	fmt.Println("func Dep_UP : Old => ", Old, Namee)
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	now := time.Now()
	setTime := now.Unix()
	coll := client.Database(dbName).Collection("tbl_hashtag")
	//ok => filter := bson.D{{"name", "Cup1"}}
	id, _ := primitive.ObjectIDFromHex(PKkey) // แปลงค่า id
	filter := bson.D{{"_id", id}}

	if Status == "ระงับการใช้งาน" {
		update := bson.D{{"$set", bson.D{{"Status", "ระงับการใช้งาน"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update Disable : ", result.MatchedCount)
	} else {
		update := bson.D{{"$set", bson.D{{"HashtagName", Namee}, {"Status", "ใช้งาน"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)

			// up tbl_Staff

			coll_staff := client.Database(dbName).Collection("tbl_hashtag")
			filterStaff := bson.D{{"Dep", Old}}
			updateStaff := bson.D{{"$set", bson.D{{"HashtagName", Dep}}}}
			resultStaff, err := coll_staff.UpdateMany(ctx, filterStaff, updateStaff)
			if err != nil {
				panic(err)
			}
		fmt.Println("Update OK : ", result.MatchedCount)

	}
	cc := gin.H{"msg": "บันทึกข้อมูลเรียบร้อยแล้ว"}
	c.JSON(200, cc)

} // end func
*/
