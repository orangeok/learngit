package logic

type User struct {
	Id          int64  `gorm:"column:id"`
	UserName    string `gorm:"column:user_name"`
	PhoneNumber string `gorm:"column:phone_number"`
	Password    int64  `gorm:"column:password"`
	Integral    int64  `gorm:"column:integral"`
	Vip         string `gorm:"column:vip"`
	ShoppingCar string `gorm:"column:shopping_car"`
}

func (User) TableName() string {
	return "user"
}

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type Sign_In struct {
	NickName    string `json:"nick_name" `
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" `
}
type UserKey struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type Commodity struct {
	Id     int64  `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	Price  int64  `gorm:"column:price"`
	Number int64  `gorm:"column:number"`
	Effect string `gorm:"column:effect"`
	Place  string `gorm:"column:place"`
}

func (Commodity) TableName() string {
	return "commodity"
}

type Orders struct {
	Id        int64  `gorm:"column:id"`
	UserName  string `gorm:"column:user_name"`
	GoodsId   string `gorm:"column:goods_id"`
	GoodsNum  int64  `gorm:"column:goods_num"`
	GoodsName string `gorm:"column:goods_name"`
}

func (Orders) TableName() string {
	return "orders"
}
