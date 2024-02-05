package purchase

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// // ดึง timestamp จากระบบแล้ว return String
func GetIimestamp(c *gin.Context) string {
	var timeassign string
	now := time.Now()
	timeassign1 := now.Unix()
	timeassign = strconv.FormatInt(timeassign1, 10)
	return timeassign
} // end func

// / แปลง date to timestamp
func TimeNow(tmp_time string) int {
	tt2, err := time.ParseInLocation("2006-01-02 15:04:05", tmp_time, time.Local)
	if err != nil {
		log.Fatal(err)
	}
	start_date := strconv.FormatInt(tt2.Unix(), 10)
	i, err := strconv.Atoi(start_date)
	return i // ค่าที่ได้คือค่า
}

// / แปลง timestamp เป็น เวลาปัจจุบัน แบบเดือนไทย
func Timestamptotime(tst int) string {
	var RetNow string
	var tmpStart string
	unixTimeUTC := time.Unix(int64(tst), 0) //gives unix time stamp in utc
	if tst > 5 {
		tmpStart = unixTimeUTC.Format("2006-01-02-15-04-05")
		RetNow2 := strings.Split(tmpStart, "-") // แยกข้อความ

		RetNow = RetNow2[2] + " " + ToThaiMon(RetNow2[1]) + " " + ToThaiYear(RetNow2[0])

		//tmpYear = RetNow2[0]

	} else {
		RetNow = "-"

	}
	return RetNow
} // end func
// / แปลง timestamp เป็น เวลาปัจจุบัน วัน-เดือน-ปี
func TimestamptoDMY(tst int) string {
	var RetNow string
	var tmpStart string
	unixTimeUTC := time.Unix(int64(tst), 0) //gives unix time stamp in utc
	if tst > 5 {
		tmpStart = unixTimeUTC.Format("2006-01-02-15-04-05")
		RetNow2 := strings.Split(tmpStart, "-") // แยกข้อความ

		//RetNow = RetNow2[2] + " " + ToThaiMon(RetNow2[1]) + " " + ToThaiYear(RetNow2[0])
		RetNow = RetNow2[2] + "-" + RetNow2[1] + "-" + RetNow2[0]

		//tmpYear = RetNow2[0]

	} else {
		RetNow = "-"

	}
	return RetNow
} // end func
// // แปลงเดือนแบบไทย
func ToThaiMon(mm string) string {
	var retnow string
	switch mm {
	case "01":
		retnow = "มกราคม"
	case "02":
		retnow = "กุมภาพันธ์"
	case "03":
		retnow = "มีนาคม"
	case "04":
		retnow = "เมษายน"
	case "05":
		retnow = "พฤษภาคม"
	case "06":
		retnow = "มิถุนายน"
	case "07":
		retnow = "กรกฎาคม"
	case "08":
		retnow = "สิงหาคม"
	case "09":
		retnow = "กันยายน"
	case "10":
		retnow = "ตุลาคม"
	case "11":
		retnow = "พฤศจิกายน"
	case "12":
		retnow = "ธันวาคม"
	default:
		retnow = "n/a"
	}
	return retnow
} // end func
// // แปลง ค.ศ. เป็น พ.ศ.
func ToThaiYear(yy string) string {
	var retnow string
	si, _ := strconv.Atoi(yy)
	si = si + 543
	retnow = strconv.Itoa(si)
	return retnow

} // end func
