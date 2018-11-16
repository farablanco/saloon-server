package models

type Points struct {
	BaseModel
	Product Product `gorm:"foreignkey:ProductRefer" db:"product_id" json:"product_id"`
	Pts     int     `db:"pts" json:"pts"`
	Price   float32 `sql:"type:decimal(18,2);" db:"price" json:"price"`
}
