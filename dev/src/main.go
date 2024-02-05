package main

/**
	108 project
	- Go Gin
	- MongoDB
**/

import (
	home "AdminProject/src/home_pk"
	product "AdminProject/src/product_pk"
	purchase "AdminProject/src/purchase_pk"
	setting "AdminProject/src/setting_pk"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// local
	//port := "8080" // ใช้ตอน deploy docker
	port := "8089"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()

	router.Use(ginsession.New())

	router.LoadHTMLGlob("templates/**/**/*.tmpl.html")
	//router.LoadHTMLGlob("templates/*/*/*.tmpl.html")
	router.Static("/static", "static")

	router.Use(gin.Logger())
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 8}) // ปิดเมื่อครบวัน 8 ชม
	router.Use(sessions.Sessions("gosession", store))
	//router.GET("/test", TestPage)
	router.GET("/welcome", ShowLogin) // login page
	router.GET("/logout", logout)     // login page

	//router.POST("/", CheckFromLogin)  // check log in CheckLogin
	//router.POST("/formAjax", ShowJsonWithData) // ajax
	//router.GET("/formAjax", ShowJsonWithData) // ajax
	// หน้าปกติ
	//router.GET("/", CheckFromLogin)  //main.go
	//router.POST("/", CheckFromLogin) // check log in CheckLogin
	/*
		// ไม่ไได้ใช้งานสำหรัลบงาน 66
		adm := router.Group("/itc")
		{
			adm.Use(AdmCheck())
			{
				adm.GET("/home", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"msg": "/itc/home"})
				})

			}

		}
	*/

	router.Use(CheckFromLogin) //CheckFromLogin router.Use(CheckStatus())
	{
		router.GET("/", MainPage)                        //main.go
		router.GET("/home", home.MainPage)               // home_pk/init.go
		router.GET("/abc/:idd", setting.URL_StaffDetail) // home_pk/init.go
		//router.POST("/home", home.MainPage)
		////// เมนูสินค้า
		prod := router.Group("/product")
		{
			prod.Use(SettingCheck())
			{

				prod.GET("/:idd", product.CheckNow)
				prod.POST("/productAjax", product.AjaxNow) //product_pk/ajax.go // Ajax

				//	prod.GET("/showStaff/:idd", product.URL_StaffDetail) //setting_pk/url.go
				//	prod.POST("/ajax", product.AjaxNow)                  //setting_pk/ajax.go

			}

		}
		/// รายละเอียด Product
		prd := router.Group("/Productdetail") //
		{
			prd.Use(PurchaseCheck())
			{

				//prd.GET("/:idd", product.URL_ProductDetail)  //  product_pk/url_product.go
				prd.POST("/show", product.URL_ProductDetail) //  product_pk/url_product.go
			}

		}
		//////////////////////////////
		////// เมนูซื้อสินค้า
		//////////////////////////////
		purch := router.Group("/purchase")
		{
			purch.Use(PurchaseCheck())
			{

				purch.GET("/:idd", purchase.CheckNow)
				//	prod.GET("/showStaff/:idd", product.URL_StaffDetail) //purchase_pk/url.go
				//	prod.POST("/ajax", product.AjaxNow)                  //purchase_pk/ajax.go
				purch.POST("/provinceAjax", purchase.AjaxNow) //purchase_pk/ajax.go // Ajax จังหวัด
				purch.POST("/purchaseAjax", purchase.AjaxNow) //purchase_pk/ajax.go // Ajax จังหวัด

			}

		}
		/// รายละเอียด Supplier
		supplier := router.Group("/Supplierdetail") //
		{
			supplier.Use(PurchaseCheck())
			{

				supplier.POST("/show", purchase.URL_SupplierDetail) //  purchase_pk/url.go

			}

		}
		//////////////////////////////
		////// เมนูตั้งค่า
		//////////////////////////////
		stg := router.Group("/setting") // สิทธิ์การใช้งานเมนู input
		{
			stg.Use(SettingCheck())
			{

				stg.GET("/:idd", setting.CheckNow)
				stg.GET("/showStaff/:idd", setting.URL_StaffDetail) //setting_pk/url.go
				stg.POST("/ajax", setting.AjaxNow)                  //setting_pk/ajax.go

			}

		}
		staff := router.Group("/staff") // กลุ่มของเจ้าหน้าที่
		{
			staff.Use(SettingCheck())
			{

				staff.GET("/:idd", setting.URL_StaffDetail)

			}

		}

		cateproduct := router.Group("/cateproduct") // หมวดสินค้า
		{
			staff.Use(SettingCheck())
			{

				cateproduct.GET("/:idd", setting.URL_CateProductDetail) // /หมวดสินค้า // setting_pk/url.go

			}

		}
		hashtag := router.Group("/hashtagdetail") // hashtag
		{
			staff.Use(SettingCheck())
			{

				hashtag.GET("/:idd", setting.URL_HashtagDetail) // /hashtag // setting_pk/url.go

			}

		}
		unitpro := router.Group("/unitproductdetail") // unitproductdetail
		{
			staff.Use(SettingCheck())
			{

				unitpro.GET("/:idd", setting.URL_ProductUnitDetail) // setting_pk/url_unitProduct.go

			}

		}
		bank := router.Group("/bankdetail") // bank
		{
			staff.Use(SettingCheck())
			{

				bank.GET("/:idd", setting.URL_BankDetail) //  setting_pk/url.go

			}

		}

	} // end CheckStatus

	router.NoRoute(func(c *gin.Context) {
		// set location
		fmt.Println("ssss : ")
		location := url.URL{Path: "/home"}
		c.Redirect(http.StatusFound, location.RequestURI())

	})
	// Group Admin

	// end New
	router.Run(":" + port)
} // end main
// /////////////////////////////////////
// / หน้าแรก
// ///////////////////////////////////////////
func MainPage(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("LOGNOW") == "LogNow" {
		cc := gin.H{"title": "หน้าหลัก1", "alert": ""}
		c.HTML(http.StatusOK, "home/home/home.tmpl.html", cc)
	}
} // end func
func CheckFromLogin(c *gin.Context) {

	session := sessions.Default(c)
	if session.Get("LOGNOW") == "LogNow" {
		fmt.Println("func CheckLogin is ok . Next return ")
		//location := url.URL{Path: "/home"}
		//c.Redirect(http.StatusFound, location.RequestURI())
		/*
			switch session.Get("LOGTYPE") {
			case "Finance":
				location := url.URL{Path: "/finance/finance/main"}
				c.Redirect(http.StatusFound, location.RequestURI())
			default:
				location := url.URL{Path: "/finance/main"}
				c.Redirect(http.StatusFound, location.RequestURI())
			}
		*/
		//return
		//	c.Next()
	} else {
		uu := c.PostForm("val_usr")
		pp := c.PostForm("val_pwd")
		loginType := c.PostForm("val_login")

		//fmt.Println("uu :", uu, " -> ", "pp : ", pp)
		if pp == "" || uu == "" {
			/// ไม่ส่งค่าอะไรมา
			/*
				cc := gin.H{"title": "เข้าสู่ระบบ"}
				c.HTML(http.StatusOK, "teml/teml/login.tmpl.html", cc)
			*/
			cc := gin.H{"title": "หน้าแรก"}

			c.HTML(http.StatusOK, "teml/teml/login.tmpl.html", cc)
		} else {
			// ตรวจสอบการ username // pass
			fmt.Println("func CheckFromlogin")
			CheckLoginNow(uu, pp, c, loginType) // main.go
		}
	}

	/**/
} // end func
// /// ตรวจสอบการ log in
func CheckLoginNow(uu string, pp string, c *gin.Context, loginType string) {
	session := sessions.Default(c)

	hasher := md5.New()
	hasher.Write([]byte(pp))
	ppencode := hex.EncodeToString(hasher.Sum(nil))
	//fmt.Println("usr : ", uu, " --pwd : ", ppencode, " --> ", loginType)
	///// connect MongoDB
	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)
	var results Staff
	coll := client.Database(dbName).Collection("tbl_Staff")
	filter := bson.M{"username": uu, "password": ppencode, "Status": "1"}
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err := coll.FindOne(ctx, filter).Decode(&results)

	if err != nil {
		logout(c)
		//	fmt.Println("Error row 177")
		//	log.Fatal(err)
	} else {
		session.Set("LOGNOW", "LogNow")
		session.Set("SID", results.ID)

		session.Save()
		var pathNow string
		pathNow = "/home"
		cc := gin.H{"EmpName": results.StaffName, "pathLink": pathNow}
		c.HTML(http.StatusOK, "teml/teml/setsessionstorage.tmpl.html", cc)
		fmt.Println("go to setsessionstorage template  ")

	}
	fmt.Println(results)
	///////////////////////////////
	/*
		if uu == "admin" && pp == "1234" {

		} else {

			//location := url.URL{Path: "/welcome"}
			// row 2
			//c.Redirect(http.StatusFound, location.RequestURI())
		}
	*/
	/*
		db := dbConn()
		//fmt.Println("uu : ", uu, " // pass : ", ppencode)

		prep, err := db.Prepare("SELECT s.idStaff, s.StaffName, s.admin_type " +
			" FROM tbl_staff s  " +
			" WHERE s.user = ?  AND s.pass = ? AND s.id_status = 1  ")

		row, err := prep.Query(uu, ppencode)
		if err != nil {
			//	log.Fatal(err)
			panic(err)
		}
		defer db.Close()
		defer row.Close()
		var staffID int
		var staffName, staffType string

		for row.Next() { // Iterate and fetch the records from result cursor

			if err := row.Scan(&staffID, &staffName, &staffType); err != nil {
				//return nil, err
			}
			// session

			session.Set("LOGNOW", "LogNow")
			session.Set("STAFFID", staffID)
			session.Set("STAFFNAME", staffName)
			session.Set("LOGTYPE", staffType)

			session.Save()
			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}
			fmt.Println("login ok ")
		} // end for
		//	fmt.Println("staffName : ", staffName)
		if staffName == "" {
			fmt.Println("staffName : is nill ")
			location := url.URL{Path: "/"}
			c.Redirect(http.StatusFound, location.RequestURI())

		} else {
			fmt.Println("go to setsessionstorage ")
			var pathNow string
			switch staffType {
			case "Finance":
				pathNow = "/finance/finance/main"
			}
			cc := gin.H{"EmpName": staffName, "pathLink": pathNow}
			c.HTML(http.StatusOK, "teml/teml/setsessionstorage.tmpl.html", cc)
			fmt.Println("go to setsessionstorage next ")
		}
	*/

} // end func

func logout(c *gin.Context) {
	fmt.Println("Func : logout")
	session := sessions.Default(c)
	session.Set("LOGNOW", "")
	session.Clear()
	session.Options(sessions.Options{MaxAge: 0}) // this sets the cookie with a MaxAge of 0
	session.Save()
	//c.Redirect(http.StatusFound, "/")
	cc := gin.H{"Status": "logout"}
	c.HTML(http.StatusOK, "teml/teml/delsessionstorage.tmpl.html", cc)

}
func ShowLogin(c *gin.Context) {

	cc := gin.H{"title": "หน้าแรก"}

	c.HTML(http.StatusOK, "teml/teml/login.tmpl.html", cc)

}

// /////////////////////////////
// / ตรวจสอบสิทธิ์การใช้งานตามเมนูต่าง ๆ
// ////////////////////////////
// / เมนู ซื้อสินค้า
func PurchaseCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ok set HandlerFunc: PurchaseCheck")

	}

} // end func
// / เมนู ตั้งค่า
func SettingCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ok set HandlerFunc: setting")

	}

} // end func
// //////////////////////////////////
// ///  MongoDB //////////////////////
// // Connect DB
func DBConn() (context.Context, *mongo.Client, string) {
	connectionURI := "mongodb://OZEadmin:MonGoOZE@127.0.0.1:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	dbName := "OZEProject"
	return ctx, client, dbName
}

/////////////////
//// ดึงข้อมูลคนมาแสดง

type Staff struct {
	ID        string `bson:"_id"`
	StaffName string `bson:"StaffName"`
	Dep       string `bson:"Dep"`
}
