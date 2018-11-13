package models

import (
	"time"
)

type Payment struct {
	BaseModel
	//Id          int    `gorm:"AUTO_INCREMENT"`
	Payment_ref_no   int     `sql:"size:10"`
	Total            float32 `sql:"type:decimal(18,2);"`
	Payment_datetime *time.Time
	Outlet           Outlet `gorm:"foreignkey:OutletRefer"`
}
