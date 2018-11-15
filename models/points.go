package models

type Points struct {
	BaseModel
	Product Product `gorm:"foreignkey:ProductRefer" db:"product_id" json:"product_id"`
	Pts     int     `db:"pts" json:"pts"`
}
