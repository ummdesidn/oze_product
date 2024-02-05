package product

/*
* function รูปภาพ
 */
import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func NewImg(c *gin.Context, idd, fileName, cc string) {
	fmt.Println("fileName : ", fileName, "--", idd)
	files, header, err := c.Request.FormFile(fileName)
	//files, err := c.FormFile("img_1")
	if err != nil {
		//panic(err)
		/*path := "./static/OZE_IMG/Product/" + idd + cc + ".jpg"
		err1 := os.Remove(path)
		if err1 != nil {
			fmt.Println(err)
			return
		}
		*/
		return

	}
	//	defer files.Close()
	fmt.Println(files)
	fmt.Println("Uploaded File: %+v\n", header.Filename)
	fmt.Println("File Size: %+v\n", header.Size)

	if files != nil {
		//dst := "./static/OZE_IMG/Product/" + idd + cc + ".jpg"
		//c.SaveUploadedFile(files, dst)
		out, err := os.Create("./static/OZE_IMG/Product/" + idd + cc + ".jpg")
		defer out.Close()
		_, err = io.Copy(out, files)
		if err != nil {
			log.Fatal(err)
		}

	}
	if err != nil {
		log.Fatal(err)
	}

}

func UpImg(c *gin.Context) {

	cc := c.PostForm("ImgNo")  // ลำดับของรูป
	idd := c.PostForm("imgID") // รหัวรูปภาพ
	fileName := idd + "_" + cc
	fmt.Println("fileName : ", fileName, "--", idd)
	files, header, err := c.Request.FormFile("img_product")
	//files, err := c.FormFile("img_product")

	if err != nil {
		panic(err)
		//return

	}
	//newFileName := uuid.New().String() + extension

	//	defer files.Close()

	fmt.Println(files)
	fmt.Println("Uploaded File: %+v\n", header.Filename)
	fmt.Println("File Size: %+v\n", header.Size)

	if files != nil {
		//dst := "./static/OZE_IMG/Product/" + idd + cc + ".jpg"
		//c.SaveUploadedFile(files, dst)
		out, err := os.Create("./static/OZE_IMG/Product/" + fileName + ".jpg")
		defer out.Close()
		_, err = io.Copy(out, files)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("Error : UpImggggggggggggggggggggggggg")
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	URL_ProductDetail(c)
} // end func

func DelImgNow(c *gin.Context) {
	// ลบรุูปสินค้า
	idd := c.PostForm("idd")
	fmt.Println("Delimg -> " + idd)
	path := "./static/OZE_IMG/Product/" + idd + ".jpg"
	err := os.Remove(path)

	if err != nil {
		//fmt.Println(err)
		//return
	}
	fmt.Println("DelImgNow() ================================================================")
	//location := url.URL{Path: "/Productdetail/show"}
	//c.Redirect(http.StatusFound, location.RequestURI())
}
