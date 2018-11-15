package models

type Payment_Item struct {
	BaseModel
	Payment  Payment `gorm:"foreignkey:PaymentRefer" db:"payment_id" json:"payment_id"`
	SubTotal float32 `sql:"type:decimal(18,2);" db:"sub_total" json:"sub_total"`
	Product  Product `gorm:"foreignkey:ProductRefer" db:"product_id" json:"product_id"`
}
