package models

type Outlet struct {
	BaseModel
	//Id          int    `gorm:"AUTO_INCREMENT"`
	Name        string `sql:"size:200" db:"name" json:"name"`
	Postal_code int    `db:"postal_code" json:"postal_code"`
}
