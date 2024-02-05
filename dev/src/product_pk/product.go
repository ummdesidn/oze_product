package product

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewProduct(c *gin.Context) {

	// 2566-7-16 ใช้งานได้ปกติ
	session := sessions.Default(c)
	//now := time.Now()
	//setTime := now.Unix()
	timee := GetIimestamp(c)

	/// 1 ข้อมูลจาก หน้าเว็บ+
	Pwd := c.PostForm("pwd")
	hasher := md5.New()
	hasher.Write([]byte(Pwd))
	enpwd := hex.EncodeToString(hasher.Sum(nil))

	//fmt.Println("pwd : ", Pwd, "--> ", session.Get("SID")) //

	ProductCate := c.PostForm("ProductCate")             // ประเภทกลุ่มสินค้า
	ProductCountStock := c.PostForm("ProductCountStock") //นับสต็อก/ไม่นับสต็อก
	productName := c.PostForm("productName")             // ชื่อสินค้า
	//IntIDproductName := c.PostForm("IntIDproductName")   // รหัสภายใน -> ยังไม่ได้ใช้งาน เนื่องจากว่ามีการสร้างอัตโนมัติเอง
	//Tmp_hid_cateid := c.PostForm("hid_cateid")
	//Tmp_hid_sup := c.PostForm("hid_sup")

	ExtIDproductName := c.PostForm("ExtIDproductName") // รหัสภายนอก
	unitBig := c.PostForm("unitBig")
	unitCount := c.PostForm("unitCount")
	unitSmall := c.PostForm("unitSmall")
	productdetail := c.PostForm("productdetail")
	productpoint := c.PostForm("productpoint")
	productcommition := c.PostForm("productcommition")
	// Supplier
	val_SuppliesName := c.PostForm("productSupplier")
	// แยกข้อความ
	tmp_SuppliesName := strings.Split(val_SuppliesName, "||")

	SuppliesName := tmp_SuppliesName[1]
	Supplies_PK := tmp_SuppliesName[0]
	// hashtag
	producthashtag := c.PostFormArray("producthashtag") // hashtag
	var Arr_Hashtag []string
	for _, hashtag := range producthashtag {
		Arr_Hashtag = append(Arr_Hashtag, hashtag)
	} // end for Hastag
	/*
		var Arr_Hashtag []string
			for _, hashtag := range producthashtag {
				Arr_Hashtag = append(Arr_Hashtag, hashtag)
			} // end for Hastag
	*/
	// ราคาสินค้า
	price_001 := c.PostForm("price_001")
	price_002 := c.PostForm("price_002")
	price_003 := c.PostForm("price_003")

	price_101 := c.PostForm("price_101")
	price_102 := c.PostForm("price_102")
	price_103 := c.PostForm("price_103")

	price_201 := c.PostForm("price_201")
	price_202 := c.PostForm("price_202")
	price_203 := c.PostForm("price_203")

	price_301 := c.PostForm("price_301")
	price_302 := c.PostForm("price_302")
	price_303 := c.PostForm("price_303")

	price_401 := c.PostForm("price_401")
	price_402 := c.PostForm("price_402")
	price_403 := c.PostForm("price_403")

	price_501 := c.PostForm("price_501")
	price_502 := c.PostForm("price_502")
	price_503 := c.PostForm("price_503")
	price_504 := c.PostForm("price_504")
	price_505 := c.PostForm("price_505")
	price_506 := c.PostForm("price_506")
	price_507 := c.PostForm("price_507")
	price_508 := c.PostForm("price_508")
	price_509 := c.PostForm("price_509")

	price_601 := c.PostForm("price_601")
	price_602 := c.PostForm("price_602")
	price_603 := c.PostForm("price_603")

	price_701 := c.PostForm("price_701")
	price_702 := c.PostForm("price_702")
	price_703 := c.PostForm("price_703")

	price_801 := c.PostForm("price_801")
	price_802 := c.PostForm("price_802")
	price_803 := c.PostForm("price_803")

	// End ราคาสินค้า
	//SuppliesName := c.PostForm("SuppliesName")
	//SuppliesContact := c.PostForm("SuppliesContact")
	//SuppliesTel := c.PostForm("SuppliesTel")
	//Supplies_id := c.PostForm("Supplies_id")

	// c.PostFormArray("xx")
	/*
		productreplace := c.PostFormArray("productreplace") // ขายทดแทน
		productselltoo := c.PostFormArray("productselltoo") // ขายประกอบ

	*/
	/// img UpImg(c *gin.Context, idd, fileName, cc string) name="img_1"
	/* ด้านล่าง ok */
	NewImg(c, timee, "img_1", "_1") // img.go
	NewImg(c, timee, "img_2", "_2")
	NewImg(c, timee, "img_3", "_3")
	NewImg(c, timee, "img_4", "_4")
	NewImg(c, timee, "img_5", "_5")
	/////
	/*
		///// save
		//// MongoDB
		/// define
		//  "xxx": xxxx, "xxx": bson.M{"xxx": xxx, "xxx": xxx, "xxxx": xxxx}, "xxx": bson.M{"xxx": xxx, "xxx": xxx, "xxxx": xxxx}
		NewProduct := bson.M{"Status": "ใช้งาน", "ProductCate": ProductCate, "ProductCountStock": ProductCountStock, "ProductName": productName, "IntIDproductName": IntIDproductName, "ExtIDproductName": ExtIDproductName, "ProductDetail": productdetail,
			"ProductUnit":     bson.M{"Big": unitBig, "Small": unitSmall, "CountTo": unitCount},
			"ProductSupplier": bson.M{"SuppliesName": SuppliesName, "SuppliesContact": SuppliesContact, "SuppliesTel": SuppliesTel, "Supplies_id": Supplies_id}, "xxx": bson.M{"xxx": xxx, "xxx": xxx, "xxxx": xxxx}}
		/// Connect
	*/
	/*
		prince {price1:{1,2,3}}
	*/

	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	tmpID := fmt.Sprintf("%v", session.Get("SID"))
	id, _ := primitive.ObjectIDFromHex(tmpID) // แปลงค่า id

	//// หารหัสกลุมสินค้า
	coll_cate := client.Database(dbName).Collection("tbl_ProductCate")
	filter_cate := bson.M{"CateName": ProductCate}
	//cursor_cate, err := coll_cate.Find(ctx, filter)
	var results_CateProduct CateProduct
	err := coll_cate.FindOne(ctx, filter_cate).Decode(&results_CateProduct)
	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	fmt.Println("results_CateProduct.GenCateID :::: ", results_CateProduct.GenCateID)

	/// หารหัสบริษัท
	coll_sid := client.Database(dbName).Collection("tbl_Supplier")
	filter_sid := bson.M{"SupName": SuppliesName}
	//cursor_cate, err := coll_cate.Find(ctx, filter)
	var results_Supplier AllSupplier
	err_sid := coll_sid.FindOne(ctx, filter_sid).Decode(&results_Supplier)
	if err_sid != nil {
		panic(err)
		log.Fatal(err_sid)
	}

	//// นับจำนวนที่ลงทะเบียน
	coll_p := client.Database(dbName).Collection("tbl_Product")
	filter_countNow := bson.M{"ProductCate": ProductCate}
	countNow, _ := coll_p.CountDocuments(ctx, filter_countNow)
	txt_runNo := fmt.Sprintf("%04d\n", int(countNow)) // ok
	txt_GenCateID := strconv.Itoa(int(results_CateProduct.GenCateID))
	txt_GenSID := strconv.Itoa(int(results_Supplier.GenSID))
	fmt.Println("txt_GenCateID => ", txt_GenCateID, "txt_GenSID => ", txt_GenSID, "txt_runNo => ", txt_runNo)
	intProductBySys := txt_GenCateID + txt_GenSID + txt_runNo

	///// เริ่มบนทึก ลง ฐานข้อมูล
	coll_check := client.Database(dbName).Collection("tbl_Staff")
	filter := bson.M{"_id": id, "password": enpwd}
	cursor, err := coll_check.Find(ctx, filter)
	defer cursor.Close(ctx)
	if err != nil {
		panic(err)

	}
	for cursor.Next(ctx) {
		///// ใช้งานได้
		coll := client.Database(dbName).Collection("tbl_Product")
		// "IntIDproductName": IntIDproductName

		doNew := bson.M{"Status": "ใช้งาน", "ProductCate": ProductCate, "ProductCountStock": ProductCountStock, "ProductName": productName, "IntIDproductName": intProductBySys, "ExtIDproductName": ExtIDproductName,
			"UnitBig": unitBig, "UnitCount": unitCount, "UnitSmall": unitSmall, "ProductDetail": productdetail, "DateIn": timee, "ProductPoint": productpoint, "ProductCommition": productcommition,
			"ProductSupplier": bson.M{"SuppliesName": SuppliesName, "Supplies_PK": Supplies_PK},
			"ProductHashtag":  Arr_Hashtag,
			"ProductPrice": bson.M{"price_001": price_001, "price_002": price_002, "price_003": price_003, "price_101": price_101, "price_102": price_102, "price_103": price_103, "price_201": price_201, "price_202": price_202, "price_203": price_203,
				"price_301": price_301, "price_302": price_302, "price_303": price_303, "price_401": price_401, "price_402": price_402, "price_403": price_403,
				"price_501": price_501, "price_502": price_502, "price_503": price_503, "price_504": price_504, "price_505": price_505, "price_506": price_506, "price_507": price_507, "price_508": price_508, "price_509": price_509,
				"price_601": price_601, "price_602": price_602, "price_603": price_603, "price_701": price_701, "price_702": price_702, "price_703": price_703, "price_801": price_801, "price_802": price_802, "price_803": price_803}}

		result, err := coll.InsertOne(ctx, doNew) // {"Contact}
		if err != nil {
			panic(err)
		}
		/*
			/// บันทึกสินค้าขายทดแทน
			coll_replace := client.Database(dbName).Collection("tbl_Product")
			filter_replace := bson.M{"DateIn": timee}
			cursor_replace, err := coll_replace.Find(ctx, filter_replace)
			if err != nil {
				panic(err)

			}
			for cursor_replace.Next(ctx) {

			}
		*/
		////// บันทึกเสร็จสิ้น

		fmt.Println("Insert OK : ", result)
		fmt.Println("Insert OK  ")

	} // end for

	/* ใช้งานได้
	///// ใช้งานได้
	coll := client.Database(dbName).Collection("tbl_Product")

	doNew := bson.M{"Status": "ใช้งาน", "ProductCate": ProductCate, "ProductCountStock": ProductCountStock, "ProductName": productName, "IntIDproductName": IntIDproductName, "ExtIDproductName": ExtIDproductName,
		"UnitBig": unitBig, "UnitCount": unitCount, "UnitSmall": unitSmall, "ProductDetail": productdetail, "DateIn": timee,
		"ProductSupplier": bson.M{"SuppliesName": SuppliesName, "Supplies_PK": Supplies_PK},
		"ProductHashtag":  Arr_Hashtag}

	result, err := coll.InsertOne(ctx, doNew) // {"Contact}
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert OK : ", result)
	*/
	////

	/*
		/// ++++ บันทึกรายละเอียด Mysql

		///// ด้านล่างของ MySql
		doProduct, err := db.Prepare("INSERT INTO tbl_product (idProduct, CountStock, idCategoryGroup, ProductName, ProductIntCode, ProductExtcode," +
			" ProductDetail, BigCount, SumCount, SmallCount, idStaff, idSupplier, status) " +

			" VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)")
		doProduct.Exec(timee, ProductCountStock, ProductCate, productName, IntIDproductName, ExtIDproductName, productdetail, unitBig, unitCount, unitSmall, session.Get("STAFFID"), productSupplier, 1)
		if err != nil {
			panic(err)
		}
	*/
	/*
		/// add สินค้าทดแทน
		for _, id_now_replace := range form_011 {
			/// วน loop บันทึก
			//fmt.Println(i, "-", id_now_replace)
			doProductReplace, err_replace := db.Prepare("INSERT INTO productreplace (idProduct , idProductReplace, idStaff) " +
				" VALUES (?,?,?)")
			doProductReplace.Exec(timee, id_now_replace, session.Get("STAFFID"))
			if err_replace != nil {
				panic(err_replace)
			}
			defer doProductReplace.Close()
		}
		/// add สินค้าขายประกอบ
		for _, id_now_other := range form_012 {
			/// วน loop บันทึก
			//fmt.Println(i, "-", id_now_other)
			doProductOther, err_replace := db.Prepare("INSERT INTO productother (idProduct , idProductOther, idStaff) " +
				" VALUES (?,?,?)")
			doProductOther.Exec(timee, id_now_other, session.Get("STAFFID"))
			if err_replace != nil {
				panic(err_replace)
			}
			defer doProductOther.Close()
		}
		/// add hashtag
		for i, id_now_hashtag := range form_013 {
			/// วน loop บันทึก
			fmt.Println(i, "-", id_now_hashtag)
			doProductHashtag, err_replace := db.Prepare("INSERT INTO producthashtag (idProduct , idHashtag, idStaff) " +
				" VALUES (?,?,?)")
			doProductHashtag.Exec(timee, id_now_hashtag, session.Get("STAFFID"))
			if err_replace != nil {
				panic(err_replace)
			}
			defer doProductHashtag.Close()
		}

		defer doProduct.Close()
		defer db.Close()
	*/
	/*
		db := dbConn()
		defer db.Close()

		do, err := db.Prepare("INSERT INTO tbl_product (idSupplier,SupplierType, SupplierID ,  Address ,ProvinceID, DistrictID, TambonID, Thaipost, idProvince, Tel, " +
			" CreditDate, TaxID, ContractPerson, ContractTel, idBank, BankAccountNo, BankAccount, CashLimit, Comment, SupplierName, idstaff, Status,SupplierPromotion ) " +
			" VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1,?)")
		do.Exec(setTime, suptype, supid, supaddress, province, amp, tam, postid, tam, suptel, credit, suptax, supcontactname, supcontacttel, idbank,
			bankaccountNo, bankaccount, cash, supcomment, supname, session.Get("STAFFID"), suppro)

		if err != nil {
			panic(err)
		}
		defer do.Close()

		for _, s := range supcates {

			//db := dbConn()
			doNow1, err11 := db.Prepare("INSERT INTO tbl_suppliercategory (idSupplier ,idCategory, idStaff, Status ) " +
				" VALUES (?,?,?,?)")
			doNow1.Exec(setTime, s, session.Get("STAFFID"), 1)

			if err11 != nil {
				panic(err11)
			}

			defer doNow1.Close()
		}
	*/

	////////////////////////////////////////////
	//	fmt.Println("Insert OK  ")
	// 3 reload
	GoNextPage("บันทึกข้อมูล Product เรียบร้อยแล้ว", "/product/all", c)

} // end func

// / ดึงข้อมูลหน้ารวม -ok
func (p AllProduct) GetProduct(c *gin.Context) (RetNow []AllProduct) {

	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Product")
	filter := bson.M{}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)
	var result AllProduct
	var results []AllProduct
	var countnow int = 1
	for cursor.Next(ctx) {
		//var result Supplyier
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		/*
			if result.SumTotal == "" {
				result.SumTotal = "0"

			}
		*/
		result.CountNow = countnow
		countnow++
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return results

} // ebd func

// / ดึงข้อมูลรายละเอียดสินค้าตาม PK
func (p AllProduct) GetProductWithPK(c *gin.Context) (RetNow []AllProduct) {
	fmt.Println("product.go => GetProductWithPK => ", c.PostForm("idd"))
	id, _ := primitive.ObjectIDFromHex(c.PostForm("idd")) // แปลงค่า id PostForm
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	coll := client.Database(dbName).Collection("tbl_Product")
	filter := bson.M{"_id": id}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	var result AllProduct
	var results []AllProduct
	var tmp_CateProduct string
	var tmp_Supplier string
	var Arr_tmp_Supplier []string
	var tmp_SupplierPK string
	var Arr_tmp_SupplierPK []string
	var ProductCate AllSupplier
	var tmp_PK primitive.ObjectID
	var ProductReplace ProductReplace
	var ProductTo ProductTo

	for cursor.Next(ctx) {
		//var result Supplyier
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		tmp_CateProduct = result.ProductCate
		tmp_PK = result.PKKey
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	//// ค้นหาชือ่บริษัท
	fmt.Println("tmp_CateProduct : ", tmp_CateProduct)
	coll_CateProduct := client.Database(dbName).Collection("tbl_Supplier")
	filter_CateProduct := bson.M{"ProductCate": tmp_CateProduct}
	cursor_CateProduct, err := coll_CateProduct.Find(ctx, filter_CateProduct)
	if err != nil {
		panic(err)
	}
	defer cursor_CateProduct.Close(ctx)
	for cursor_CateProduct.Next(ctx) {

		if err := cursor_CateProduct.Decode(&ProductCate); err != nil {
			log.Fatal(err)
		}
		tmp_Supplier = ProductCate.SupName
		Arr_tmp_Supplier = append(Arr_tmp_Supplier, tmp_Supplier)
		tmp_SupplierPK = ProductCate.PKKey.String()
		Arr_tmp_SupplierPK = append(Arr_tmp_SupplierPK, tmp_SupplierPK)

	} // end for
	result.ProductSupplier.SuppliersProductCateArr = Arr_tmp_Supplier
	result.ProductSupplier.SuppliersPKArr = Arr_tmp_SupplierPK // _id
	results = append(results, result)
	/// รายการสินค้าทดแทน
	/// เอา ProductCate = result.ProductCate
	var Tmp_ProductReplaceName string
	var Arr_ProductReplaceName []string
	var Tmp_ProductReplacePK string
	var Arr_ProductReplacePK []string
	coll_ReplaceProduct := client.Database(dbName).Collection("tbl_Product")
	filter_ReplaceProduct := bson.D{{"$and",
		bson.A{
			bson.M{"ProductCate": tmp_CateProduct},

			bson.D{{"_id", bson.D{{"$ne", tmp_PK}}}},
			bson.M{"Status": "ใช้งาน"},
		},
	},
	} // end filter filter_ReplaceProduct
	//fmt.Println("Tmp_ProductReplace -- > Start ")
	cursor_ReplaceProduct, err := coll_ReplaceProduct.Find(ctx, filter_ReplaceProduct)
	defer cursor_ReplaceProduct.Close(ctx)

	for cursor_ReplaceProduct.Next(ctx) {

		if err := cursor_ReplaceProduct.Decode(&ProductReplace); err != nil {
			log.Fatal(err)
		}
		Tmp_ProductReplaceName = ProductReplace.ProductName
		Arr_ProductReplaceName = append(Arr_ProductReplaceName, Tmp_ProductReplaceName)
		Tmp_ProductReplacePK = ProductReplace.PKKey.String()
		Arr_ProductReplacePK = append(Arr_ProductReplacePK, Tmp_ProductReplacePK)

		//fmt.Println("Tmp_ProductReplace -- > ", Tmp_ProductReplaceName)

	} // end for
	result.ProductReplace.ProductReplaceNameArr = Arr_ProductReplaceName
	result.ProductReplace.ProductReplacePKArr = Arr_ProductReplacePK
	results = append(results, result)
	//fmt.Println("result.ProductReplace.ProductReplaceNameArr -- > ", result.ProductReplace.ProductReplaceNameArr, " === ", result.ProductReplace.ProductReplacePKArr)
	//fmt.Println("Tmp_ProductReplace -- > End ")
	//// end Product Replace
	/// รายการสินค้าประกอบ

	var Tmp_ProductToName string
	var Arr_ProductToName []string
	var Tmp_ProductToPK string
	var Arr_ProductToPK []string
	coll_ToProduct := client.Database(dbName).Collection("tbl_Product")
	filter_ToProduct := bson.D{{"$and",
		bson.A{
			bson.M{"ProductCate": bson.D{{"$ne", tmp_CateProduct}}},

			bson.M{"Status": "ใช้งาน"},
		},
	},
	} // end filter filter_ReplaceProduct
	fmt.Println("Tmp_ProductTo -- > Start ")
	cursor_ToProduct, err := coll_ToProduct.Find(ctx, filter_ToProduct)
	defer cursor_ToProduct.Close(ctx)

	for cursor_ToProduct.Next(ctx) {

		if err := cursor_ToProduct.Decode(&ProductTo); err != nil {
			log.Fatal(err)
		}
		Tmp_ProductToName = ProductTo.ProductName
		Arr_ProductToName = append(Arr_ProductToName, Tmp_ProductToName)
		Tmp_ProductToPK = ProductTo.PKKey.String()
		Arr_ProductToPK = append(Arr_ProductToPK, Tmp_ProductToPK)

		//fmt.Println("Tmp_ProductTo -- > ", Tmp_ProductReplaceName)

	} // end for

	result.ProductTo.ProductToNameArr = Arr_ProductToName
	result.ProductTo.ProductToPKArr = Arr_ProductToPK
	results = append(results, result)
	fmt.Println("result.ProductTo.ProductToNameArr -- > ", result.ProductTo.ProductToNameArr, " === ", result.ProductTo.ProductToPKArr)
	fmt.Println("Tmp_ProductTo -- > End ")
	//// end รายการสินค้าประกอบ

	return results
} // ebd func

// ///////////////////////////////////////////////////////////////////////
// ปรับปรุงรายละเอียดสินค้า
func UpProduct(c *gin.Context) {
	fmt.Println("UpProduct")
	/// 1 ข้อมูลจาก หน้าเว็บ+
	PKkey := c.PostForm("pk")                            // PK
	Status := c.PostForm("Status")                       // สถานะ
	ProductCate := c.PostForm("ProductCate")             // ประเภทกลุ่มสินค้า
	ProductCountStock := c.PostForm("ProductCountStock") //นับสต็อก/ไม่นับสต็อก
	productName := c.PostForm("productName")             // ชื่อสินค้า
	//IntIDproductName := c.PostForm("IntIDproductName")   // รหัสภายใน
	ExtIDproductName := c.PostForm("ExtIDproductName") // รหัสภายนอก
	unitBig := c.PostForm("unitBig")
	unitCount := c.PostForm("unitCount")
	unitSmall := c.PostForm("unitSmall")
	productdetail := c.PostForm("productdetail")
	//imgID := c.PostForm("imgID") // รหัวรูปภาพ
	// Supplier
	val_SuppliesName := c.PostForm("productSupplier")
	tmp_SuppliesName := strings.Split(val_SuppliesName, "||")
	SuppliesName := tmp_SuppliesName[1]
	Supplies_PK := tmp_SuppliesName[0]
	/// hashtag
	producthashtag := c.PostFormArray("producthashtag") // hashtag
	/// 2 connect
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	now := time.Now()
	setTime := now.Unix()
	coll := client.Database(dbName).Collection("tbl_Product")
	//ok => filter := bson.D{{"name", "Cup1"}}
	//id, _ := primitive.ObjectIDFromHex("65313f3fb4fbdbe392b297ba") // แปลงค่า id
	id, _ := primitive.ObjectIDFromHex(PKkey)

	filter := bson.D{{"_id", id}}
	if Status == "ระงับการใช้งาน" {
		update := bson.D{{"$set", bson.D{{"Status", "ระงับการใช้งาน"}, {"Checkpoint", setTime}}}}
		result, _ := coll.UpdateOne(ctx, filter, update)
		fmt.Println("Update Disable : ", result.MatchedCount)
	} else {
		/*	update := bson.D{{"$set", bson.D{{"Status", "ใช้งาน"}, {"Address", supaddress}, {"SupID", supid}, {"SupType", suptype}, {"AmpurID", amp}, {"PostID", postid}, {"ProvinceID", province}, {"SupContactName", supcontactname}, {"SupContactTel", supcontacttel}, {
				"SupName", supname}, {"SupTaxID", suptax}, {"TambonID", tam}, {"Comment", supcomment}, {"CheckPoint", setTime}, {"SupPromotion", suppro}, {
				"DateIn", setTime}, {"SupTel", suptel}, {"Bank", bson.D{{"BankName", bankbrand}, {"BankNo", bankno}, {"BookBank", bankaccount}}}, {"Credit", bson.D{{"Date", credit}, {"Cash", cash}}}, {"ProductCate", ar2}}}}
			result, _ := coll.UpdateOne(ctx, filter, update)
		*/

		// hashtag
		var Arr_Hashtag []string
		for _, hashtag := range producthashtag {
			Arr_Hashtag = append(Arr_Hashtag, hashtag)
		} // end for Hastag

		doUp := bson.D{{"$set", bson.D{{"Status", "ใช้งาน"}, {"ProductCate", ProductCate}, {"ProductCountStock", ProductCountStock}, {"ProductName", productName},
			{"ExtIDproductName", ExtIDproductName}, {"UnitBig", unitBig}, {"UnitCount", unitCount}, {"UnitSmall", unitSmall}, {"ProductDetail", productdetail}, {"Checkpoint", setTime},
			{"ProductSupplier", bson.D{{"SuppliesName", SuppliesName}, {"Supplies_PK", Supplies_PK}}},
			{"ProductHashtag", Arr_Hashtag}}}}
		result, _ := coll.UpdateOne(ctx, filter, doUp)

		fmt.Println("Arr_Hashtag :: ", Arr_Hashtag)

		fmt.Println("Update OK : ", result.MatchedCount)
		/// img
		//	UpImg(c, imgID, "img_1", "_1")
		//	UpImg(c, imgID, "img_2", "_2")
		//	UpImg(c, imgID, "img_3", "_3")
		//	UpImg(c, imgID, "img_4", "_4")
		//	UpImg(c, imgID, "img_5", "_5")
		// upPrince
		UpPriceProduct(c) // product.go

	}
	/// 3 reload
	GoNextPage("บันทึกข้อมูลรายละเอียดสินค้า เรียบร้อยแล้ว", "/product/all", c) // alert.go
	/*
		session := sessions.Default(c)
		//now := time.Now()
		//setTime := now.Unix()
		/// 1 ข้อมูลจาก หน้าเว็บ+

		PKkey := c.PostForm("pk")
		Status := c.PostForm("Status")
		ProductCate := c.PostForm("ProductCate")
		ProductCountStock := c.PostForm("ProductCountStock")
		productName := c.PostForm("productName")
		IntIDproductName := c.PostForm("IntIDproductName")
		ExtIDproductName := c.PostForm("ExtIDproductName")
		unitBig := c.PostForm("unitBig")
		unitCount := c.PostForm("unitCount")
		unitSmall := c.PostForm("unitSmall")
		productSupplier := c.PostForm("productSupplier")
		productdetail := c.PostForm("productdetail")

		///////
		db := DBConn()
		defer db.Close()
		if Status == "0" {
			doUp, err_doUp := db.Prepare("UPDATE  tbl_product  SET  status= 0 , idstaff=?    WHERE idProduct = ? ")
			doUp.Exec(session.Get("STAFFID"), PKkey)
			if err_doUp != nil {
				panic(err_doUp)
			}
			defer doUp.Close()

		} else {
			fmt.Println("----- Update Product ")

			doUp, err_doUp := db.Prepare("UPDATE  tbl_product  SET  status= 1 , " +
				" idCategoryGroup=?, CountStock=?, ProductName=?, ProductIntCode=?, ProductExtcode=?, " +
				" BigCount=?, SumCount=?, SmallCount=?, idSupplier=?, ProductDetail=?, idStaff=? " +
				" WHERE idProduct = ? ")

			doUp.Exec(ProductCate, ProductCountStock, productName, IntIDproductName, ExtIDproductName, unitBig, unitCount, unitSmall, productSupplier, productdetail, session.Get("STAFFID"), PKkey)
			if err_doUp != nil {
				panic(err_doUp)
			}
			/// img
			UpImg(c, PKkey, "img_1", "_1")
			UpImg(c, PKkey, "img_2", "_2")
			UpImg(c, PKkey, "img_3", "_3")
			UpImg(c, PKkey, "img_4", "_4")
			UpImg(c, PKkey, "img_5", "_5")

			///// ขาดสินค้า ทดแทน ขายประกอบ Hashtag
			/// เปลี่ยนแปลงหมวดสินค้า

			defer doUp.Close()

		}


	*/
} // end func
// ////////////////////////////
// ปรับปรุงราคาสินค้า
func UpPriceProduct(c *gin.Context) {
	fmt.Println("UpPriceProduct")
	/// 1 ข้อมูลจาก หน้าเว็บ+
	PKkey := c.PostForm("pk") // PK
	// ราคาสินค้า
	price_001 := c.PostForm("price_001")
	price_002 := c.PostForm("price_002")
	price_003 := c.PostForm("price_003")

	price_101 := c.PostForm("price_101")
	price_102 := c.PostForm("price_102")
	price_103 := c.PostForm("price_103")

	price_201 := c.PostForm("price_201")
	price_202 := c.PostForm("price_202")
	price_203 := c.PostForm("price_203")

	price_301 := c.PostForm("price_301")
	price_302 := c.PostForm("price_302")
	price_303 := c.PostForm("price_303")

	price_401 := c.PostForm("price_401")
	price_402 := c.PostForm("price_402")
	price_403 := c.PostForm("price_403")

	price_501 := c.PostForm("price_501")
	price_502 := c.PostForm("price_502")
	price_503 := c.PostForm("price_503")
	price_504 := c.PostForm("price_504")
	price_505 := c.PostForm("price_505")
	price_506 := c.PostForm("price_506")
	price_507 := c.PostForm("price_507")
	price_508 := c.PostForm("price_508")
	price_509 := c.PostForm("price_509")

	price_601 := c.PostForm("price_601")
	price_602 := c.PostForm("price_602")
	price_603 := c.PostForm("price_603")

	price_701 := c.PostForm("price_701")
	price_702 := c.PostForm("price_702")
	price_703 := c.PostForm("price_703")

	price_801 := c.PostForm("price_801")
	price_802 := c.PostForm("price_802")
	price_803 := c.PostForm("price_803")

	/// 2 connect
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	//now := time.Now()
	//setTime := now.Unix()
	coll := client.Database(dbName).Collection("tbl_Product")

	id, _ := primitive.ObjectIDFromHex(PKkey)

	filter := bson.D{{"_id", id}}

	doUp := bson.D{{"$set", bson.D{{"ProductPrice", bson.D{{"price_001", price_001}, {"price_002", price_002}, {"price_003", price_003}, {"price_101", price_101}, {"price_102", price_102}, {"price_103", price_103}, {"price_201", price_201}, {"price_202", price_202}, {"price_203", price_203}, {"price_301", price_301},
		{"price_302", price_302}, {"price_303", price_303}, {"price_401", price_401}, {"price_402", price_402}, {"price_403", price_403}, {"price_501", price_501}, {"price_502", price_502}, {"price_503", price_503}, {"price_504", price_504}, {"price_505", price_505},
		{"price_506", price_506}, {"price_507", price_507}, {"price_508", price_508}, {"price_509", price_509}, {"price_601", price_601}, {"price_602", price_602}, {"price_603", price_603}, {"price_701", price_701}, {"price_702", price_702}, {"price_703", price_703},
		{"price_801", price_801}, {"price_802", price_802}, {"price_803", price_803}}}}}}

	result, _ := coll.UpdateOne(ctx, filter, doUp)
	fmt.Println("Update OK : ", result.MatchedCount)
	//// 3 update รายละเอียด
	///// กำลังทำ
	/// 4 reload
	GoNextPage("บันทึกข้อมูลรายละเอียดสินค้า เรียบร้อยแล้ว", "/product/all", c) // alert.go

} // end func

/////////////////////////////
//// ด้านล่างไม่ได้ใช่้งาน
/////////////////////////////////////////////
/*



// / ดึงข้อมูล Supplier ตามค่า PPkey
func (p AllSupplier) GetSupplierByID(c *gin.Context, idd string) (RetNow []AllSupplier) {
	db := dbConn()
	prep, err1 := db.Prepare("SELECT COALESCE(sup.idSupplier ,''), COALESCE(sup.Status,''), COALESCE(sup.SupplierType,''), " +
		" COALESCE(sup.SupplierID,''), COALESCE(sup.SupplierName,''), COALESCE(sup.TaxID,''), COALESCE(sup.Address,''), " +
		" COALESCE(sup.Thaipost,''), COALESCE(sup.Tel,''), COALESCE(sup.ContractPerson,''), COALESCE(sup.ContractTel,''), " +
		" COALESCE(sup.idBank,''), COALESCE(sup.BankAccountNo,''), COALESCE(sup.BankAccount,''), COALESCE(sup.CreditDate,''), " +
		" COALESCE(sup.CashLimit,''), COALESCE(sup.SupplierPromotion,''), COALESCE(sup.Comment,''), " +
		" COALESCE(sup.ProvinceID,''),COALESCE(prov.DistrictThaiShort,''),COALESCE(prov.TambonThaiShort,''), " +
		" COALESCE(sup.DistrictID,''),COALESCE(sup.TambonID,''), " +
		" COALESCE(supCate.idCategory,'')" +
		" FROM tbl_supplier sup   " +
		" LEFT JOIN tbl_province prov ON prov.idTambon = sup.TambonID " +
		" LEFT JOIN tbl_suppliercategory supCate ON supCate.idSupplier = sup.idSupplier " +

		" WHERE (sup.idSupplier = ? )")
	// ชาด  หมวดสินค้า
	row, err1 := prep.Query(idd)
	if err1 != nil {
		panic(err1)
	}
	var count int
	var inTmpCate string
	var inArrTmpCate []string
	for row.Next() {
		var tmp AllSupplier
		count += 1
		tmp.Count = count

		if err := row.Scan(&tmp.PKKey, &tmp.Status, &tmp.SupplyierType, &tmp.SupplyierID, &tmp.SupplyierName, &tmp.SupplyierTax, &tmp.SupplyierAddress,
			&tmp.SupplyierPostID, &tmp.SupplyierTel, &tmp.SupplyierContact, &tmp.SupplyierContactTel,
			&tmp.SupplyierBank, &tmp.SupplyierAccountNo, &tmp.SupplyierAccountBank, &tmp.SupplyierCredit,
			&tmp.SupplyierCashLimit, &tmp.SupplyierPromotion, &tmp.SupplyierComment,
			&tmp.SupplyierProvinceID, &tmp.SupplyierDistrictName_txt, &tmp.SupplyierDistrictName_txt,
			&tmp.SupplyierDistrictID, &tmp.SupplyierTambonID,
			&inTmpCate); err != nil {
			panic(err)
		}
		/// หาค่าอำเภอ
		prep_Dist, _ := db.Prepare("SELECT DISTINCT COALESCE(prov.DistrictID,0),COALESCE(prov.DistrictThaiShort,'') " +
			" FROM tbl_province prov " +
			" WHERE (prov.ProvinceID = ? ) ")
		row_Dist, err1 := prep_Dist.Query(tmp.SupplyierProvinceID)
		if err1 != nil {
			panic(err1)
		}
		var tmpDistID int
		var tmpDisttxt string
		var tmpArrDistID []int
		var tmpArrDisttxt []string
		for row_Dist.Next() {

			if err := row_Dist.Scan(&tmpDistID, &tmpDisttxt); err != nil {
				panic(err)
			}
			tmpArrDistID = append(tmpArrDistID, tmpDistID)
			tmpArrDisttxt = append(tmpArrDisttxt, tmpDisttxt)
			tmp.ArrSupplyierDistrictID = tmpArrDistID
			tmp.ArrSupplyierDistrictName_txt = tmpArrDisttxt

		}
		defer row_Dist.Close()
		/// หาค่าตำบล
		prep_tambont, _ := db.Prepare("SELECT DISTINCT COALESCE(prov.idTambon,0),COALESCE(prov.TambonThaiShort,'') " +
			" FROM tbl_province prov " +
			" WHERE (prov.DistrictID = ? ) ")
		row_tambon, err1 := prep_tambont.Query(tmp.SupplyierDistrictID)
		if err1 != nil {
			panic(err1)
		}
		var tmpTamID int
		var tmpTamtxt string
		var tmpArrTamID []int
		var tmpArrTamtxt []string
		for row_tambon.Next() {

			if err := row_tambon.Scan(&tmpTamID, &tmpTamtxt); err != nil {
				panic(err)
			}
			tmpArrTamID = append(tmpArrTamID, tmpTamID)
			tmpArrTamtxt = append(tmpArrTamtxt, tmpTamtxt)
			tmp.ArrSupplyierTambonID = tmpArrTamID
			tmp.ArrSupplyierTambonName_txt = tmpArrTamtxt

		}
		defer row_Dist.Close()

		/// หาค่าหมวดสินค้าที่ใช้งาน
		inArrTmpCate = append(inArrTmpCate, inTmpCate)
		tmp.TmpCate = inArrTmpCate
		///  end
		RetNow = append(RetNow, tmp)
	}
	defer row.Close()
	defer db.Close()
	return RetNow

} // end func
// / ปรัลปรุงรายละเอียด Supplier


*/
