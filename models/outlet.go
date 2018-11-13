package models

type Outlet struct {
	BaseModel
	//Id          int    `gorm:"AUTO_INCREMENT"`
	Name        string `sql:"size:200"`
	Postal_code int
}
