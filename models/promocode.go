package models

type Promocode struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Code      string `json:"code"`
	FromPrice string `json:"from_price"`
	Discount  string `json:"discount"`
}
