package models

type Points struct {
	BaseModel
	//Id          int    `gorm:"AUTO_INCREMENT"`
	Product Product `gorm:"foreignkey:ProductRefer"`
	Pts     int
}
