package models

import (
	"time"
)

type Payment struct {
	BaseModel
	Payment_ref_no   int        `sql:"size:10" db:"payment_ref_no" json:"payment_ref_no"`
	Total            float32    `sql:"type:decimal(18,2);" db:"total" json:"total"`
	Payment_datetime *time.Time `db:"payment_datetime" json:"payment_datetime"`
	Outlet           Outlet     `gorm:"foreignkey:OutletRefer" db:"outlet_id" json:"outlet_id"`
	Discount         float32    `sql:"type:double;" db:"discount" json:"discount"`
}
