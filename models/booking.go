// author : Kenneth Phang
package models

import "time"

type Booking struct {
	Id              int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt       time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deletedAt"`
	Hairdresser     int64      `db:"hairdresser_id" json:"hairDresserId"`
	BookingDatetime *time.Time `db:"booking_datetime" json:"bookingDatetime"`
	Outlet          int64      `db:"outlet_id" json:"outletId"`
	Customer        int64      `db:"customer_id" json:"customerId"`
}
