package models

import (
	"time"
)

type Payment struct {
	Id              int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt       time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deletedAt"`
	PaymentRefNo    string     `sql:"size:10" db:"payment_ref_no" json:"paymentRefNo"`
	Total           float32    `sql:"type:decimal(18,2);" db:"total" json:"total"`
	PaymentDatetime *time.Time `db:"payment_datetime" json:"paymentDatetime"`
	Outlet          Outlet     `gorm:"foreignkey:OutletRefer" db:"outlet_id" json:"outletId"`
	Discount        float32    `sql:"type:double;" db:"discount" json:"discount"`
	PaymentMode     int        `sql:"type:double;" db:"payment_mode" json:"paymentMode"`
	User            User       `gorm:"foreignkey:UserRefer" db:"user_id" json:"user"`
	Gst             float32    `sql:"type:double;" db:"gst" json:"gst"`
}
