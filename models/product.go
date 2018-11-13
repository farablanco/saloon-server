package models

type Product struct {
	BaseModel
	//Id          int    `gorm:"AUTO_INCREMENT"`
	Name   string  `sql:"size:200"`
	price  float32 `sql:"type:decimal(18,2);"`
	status string
}
