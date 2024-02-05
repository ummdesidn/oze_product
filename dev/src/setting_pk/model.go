package setting

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	PKKey     primitive.ObjectID `bson:"_id"`
	StaffName string             `bson:"StaffName"`
	Dep       string             `bson:"Dep"`

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
	PKKey      primitive.ObjectID `bson:"_id"`
	TXT_ID     string
	DepName    string `bson:"DepName"`
	Status     string `bson:"Status"`
	Checkpoint int64  `bson:"Checkpoint"`
}

type CateProduct struct {
	PKKey    primitive.ObjectID `bson:"_id"`
	CateName string             `bson:"CateName"`
	Status   string             `bson:"Status"`
}

type Hashtag struct {
	PKKey       primitive.ObjectID `bson:"_id"`
	HashtagName string             `bson:"HashtagName"`
	Status      string             `bson:"Status"`
}

type Bank struct {
	PKKey      primitive.ObjectID `bson:"_id"`
	BankName   string             `bson:"BankName"`
	Status     string             `bson:"Status"`
	Checkpoint int64              `bson:"Checkpoint"`
}
type UnitProduct struct {
	PKKey    primitive.ObjectID `bson:"_id"`
	UnitName string             `bson:"UnitName"`
	Status   string             `bson:"Status"`
}
