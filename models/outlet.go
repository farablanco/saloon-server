package models

import "time"

type Outlet struct {
	Id         int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt  time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deletedAt"`
	Name       string     `sql:"size:200" db:"name" json:"name"`
	PostalCode int        `db:"postal_code" json:"postal_code"`
	ContactNo  string     `sql:"size:45" db:"contact_no" json:"contactNo"`
	Email      string     `sql:"size:100" db:"email" json:"email"`
}
