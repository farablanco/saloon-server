package models

type Payment_Item struct {
	BaseModel
	//Id          int    `gorm:"AUTO_INCREMENT"`
	Payment  Payment `gorm:"foreignkey:PaymentRefer"`
	SubTotal float32 `sql:"type:decimal(18,2);"`
	Discount float32 `sql:"type:double;"`
}
