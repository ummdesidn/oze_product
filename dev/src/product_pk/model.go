package product

import "go.mongodb.org/mongo-driver/bson/primitive"

/////// ด้านล่าง ลบเมื่อเสร้จแล้ว

type Employee struct {
	PKKey     primitive.ObjectID `bson:"_id"`
	StaffName string             `bson:"StaffName"`
	Dep       Dep                `bson:"Dep"`

	JobLevel string `bson:"JobLevel"`

	Username string `bson:"username"`
	Status   string `bson:"Status"`

	CheckPoint int64   `bson:"CheckPoint"`
	Contacts   Contact `bson:"Contact"`
}
type Dep struct {
	PKKey primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"Name"`
}
type Contact struct {
	Tel    string `bson:"Tel"`
	Mobile string `bson:"Mobile"`
	Email  string `bson:"Email"`
}

type Depp struct {
	PKKey   primitive.ObjectID `bson:"_id"`
	TXT_ID  string
	DepName string `bson:"DepName"`
	Status  string `bson:"Status"`
}

type CateProduct struct {
	PKKey     primitive.ObjectID `bson:"_id"`
	CateName  string             `bson:"CateName"`
	Status    string             `bson:"Status"`
	GenCateID int64              `bson:"GenCateID"`
}
type AllSupplier struct {
	PKKey primitive.ObjectID `bson:"_id"`
	//SupPK       string             `bson:"_id"`
	SupName     string   `bson:"SupName"`
	ProductCate []string `bson:"ProductCate"`
	GenSID      int64    `bson:"GenSID"` /// ค่าที่ระบบ gen ขึ้นมา int64

	Status         string `bson:"Status"`
	Address        string `bson:"Address"`
	SupID          string `bson:"SupID"`
	SupType        string `bson:"SupType"`
	AmpurID        string `bson:"AmpurID"`
	PostID         string `bson:"PostID"`
	ProvinceID     string `bson:"ProvinceID"`
	SupContactName string `bson:"SupContactName"`
	SupContactTel  string `bson:"SupContactTel"`
	SupTaxID       string `bson:"SupTaxID"`
	TambonID       string `bson:"TambonID"`
	Comment        string `bson:"Comment"`
	LastPO         string `bson:"LastPO"`
	SumYear        string `bson:"SumYear"`
	SupTel         string `bson:"SupTel"`
}

/// หนาวนนับสินค้า
type AllUnitProduct struct {
	PKKey     string `json:"PKKey"`
	UnitName  string `json:"UnitName"`
	Status    string `json:"Status"`
	TxtStatus string `json:"TxtStatus"` // txt สถานะ
}
type AllProduct struct {
	PKKey             primitive.ObjectID `bson:"_id"`
	CountNow          int                `bson:"CountNow"`
	ProductCate       string             `bson:"ProductCate"`
	TxtProductCate    string             `bson:"TxtProductCate"`
	ProductCountStock string             `bson:"ProductCountStock"`
	ProductName       string             `bson:"ProductName"`
	IntIDproductName  string             `bson:"IntIDproductName"`
	ExtIDproductName  string             `bson:"ExtIDproductName"`
	UnitBig           string             `bson:"UnitBig"`
	UnitCount         string             `bson:"UnitCount"`
	UnitSmall         string             `bson:"UnitSmall"`
	ProductSupplier   SupplierDetail     `bson:"ProductSupplier"`
	Productdetail     string             `bson:"productdetail"`
	SumTotal          string             `bson:"SumTotal"` // จำนวนของสินค้าที่มี
	Status            string             `bson:"Status"`   // จำนวนของสินค้าที่มี
	DateIn            string             `bson:"DateIn"`
	ProductHashtag    []string           `bson:"ProductHashtag"`
	ProductPrice      AllPrice           `bson:"ProductPrice"` /// ราคาขาย
	ProductPoint      string             `bson:"ProductPoint"`
	ProductCommition  string             `bson:"ProductCommition"`
}
type SupplierDetail struct {
	PKKey                   string   `bson:"Supplies_PK"`
	SuppliesName            string   `bson:"SuppliesName"`
	SuppliersProductCateArr []string `bson:"SuppliersProductCateArr"` // เก็บค่าชื่อ ผู็จัดจำหน่าย
	SuppliersPKArr          []string `bson:"SuppliersPKArr"`          // เก็บค่าชื่อ _id ผู็จัดจำหน่าย
}
type AllHashtag struct {
	HashtagName string `bson:"HashtagName"`
	Status      string `bson:"Status"`
}

type AllPrice struct {
	Price_001 string `bson:"price_001"`
	Price_002 string `bson:"price_002"`
	Price_003 string `bson:"price_003"`

	Price_101 string `bson:"price_101"`
	Price_102 string `bson:"price_102"`
	Price_103 string `bson:"price_103"`

	Price_201 string `bson:"price_201"`
	Price_202 string `bson:"price_202"`
	Price_203 string `bson:"price_203"`

	Price_301 string `bson:"price_301"`
	Price_302 string `bson:"price_302"`
	Price_303 string `bson:"price_303"`

	Price_401 string `bson:"price_401"`
	Price_402 string `bson:"price_402"`
	Price_403 string `bson:"price_403"`

	Price_501 string `bson:"price_501"`
	Price_502 string `bson:"price_502"`
	Price_503 string `bson:"price_503"`
	Price_504 string `bson:"price_504"`
	Price_505 string `bson:"price_505"`
	Price_506 string `bson:"price_506"`
	Price_507 string `bson:"price_507"`
	Price_508 string `bson:"price_508"`
	Price_509 string `bson:"price_509"`

	Price_601 string `bson:"price_601"`
	Price_602 string `bson:"price_602"`
	Price_603 string `bson:"price_603"`

	Price_701 string `bson:"price_701"`
	Price_702 string `bson:"price_702"`
	Price_703 string `bson:"price_703"`

	Price_801 string `bson:"price_801"`
	Price_802 string `bson:"price_802"`
	Price_803 string `bson:"price_803"`
}

/////////////////////////////////////////////////
//////// mysql +++++++++ๅ
/// รายละเอียดสินค้า
type AllProductDeatil struct {
	CountNow          int    `bson:"CountNow"`
	PKKey             string `bson:"PKKey"`
	ProductCate       string `bson:"ProductCate"`
	TxtProductCate    string `bson:"TxtProductCate"`
	ProductCountStock string `bson:"ProductCountStock"`
	ProductName       string `bson:"ProductName"`
	IntIDproductName  string `bson:"IntIDproductName"`
	ExtIDproductName  string `bson:"ExtIDproductName"`
	UnitBig           string `bson:"unitBig"`
	UnitCount         string `bson:"unitCount"`
	UnitSmall         string `bson:"unitSmall"`
	ProductSupplier   string `bson:"productSupplier"`
	Productdetail     string `bson:"productdetail"`
	SumTotal          string `bson:"SumTotal"` // จำนวนของสินค้าที่มี
	Status            string `bson:"Status"`   // จำนวนของสินค้าที่มี
}

/*
// หมวดสินค้า
type ALLCateProduct struct {
	PKKey     string `json:"PKKey"`
	CateName  string `json:"CateName"`
	Status    string `json:"Status"`
	TxtStatus string `json:"TxtStatus"`
}*/
