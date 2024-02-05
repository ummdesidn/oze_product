package purchase

import "go.mongodb.org/mongo-driver/bson/primitive"

type Supplyier struct {
	PKKey               primitive.ObjectID `bson:"_id"`
	SupplyierType       string             `bson:"SupType"`
	SupplyierTax        string             `bson:"SupTaxID"`
	SupplyierID         string             `bson:"SupID"`
	Suptitlename        string             `bson:"Suptitlename"`
	SupplyierName       string             `bson:"SupName"`
	SupplyierTel        string             `bson:"SupTel"`
	SupplyierCredit     Credit             `bson:"Credit"`
	SupplyierContact    string             `bson:"SupContactName"`
	SupplyierContactTel string             `bson:"SupContactTel"`
	Status              string             `bson:"Status"`
	Bank                Bank               `bson:"Bank"`
	Count               int
	Comment             string `bson:"Comment"`
	DateIn              int64  `bson:"DateIn"`
	RegisterDateTXT     string `bson:"RegisterDateTXT"`
	ProvinceID          string `bson:"ProvinceID"`
	DistrictName        string `bson:"AmpurID"`  // amp
	TambonName          string `bson:"TambonID"` // tam
	PostID              string `bson:"PostID"`

	SupplyierCate []string `bson:"ProductCate"` // หมวดสินค้า
	//CheckPoint int64 `bson:"CheckPoint"`

}

/// ข้อมุลธนาคาร
type Bank struct {
	PKKey    primitive.ObjectID `bson:"_id"`
	BankName string             `bson:"BankName"`
	//ShowBankName string             `bson:"Name"`
	BankNo   string `bson:"BankNo"`
	BookBank string `bson:"BookBank"`
}

type Province struct {
	ID           primitive.ObjectID `bson:"_id"`
	ProvinceThai string             `bson:"ProvinceThai"`
	DistrictThai string             `bson:"DistrictThai"`
}

/// ข้อมูล credit
type Credit struct {
	CreditDate string `bson:"Date"`
	CreditCash string `bson:"Cash"`
}

// หมวดสินค้า
type CateProduct struct {
	PKKey       primitive.ObjectID `bson:"_id"`
	CateName    string             `bson:"CateName"`
	Status      string             `bson:"Status"`
	Test_Status int                `bson:"Test_Status"`
}

/////// ด้านล่าง ลบเมื่อเสร้จแล้ว
/*


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



/// สถานที่



*/
