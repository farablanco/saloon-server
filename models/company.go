// author : Kenneth Phang
package models

import "time"

type Company struct {
	Id               int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt        time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt        *time.Time `db:"deleted_at" json:"deletedAt"`
	Name             string     `sql:"size:200" db:"name" json:"name"`
	RegistrationCode string     `sql:"size:200" db:"rcbNo" json:"rcbNo"`
	ContactNo        string     `sql:"size:45" db:"contact_no" json:"contactNo"`
	Email            string     `sql:"size:100" db:"email" json:"email"`
	CompanyOwner     string     `sql:"size:100" db:"supervisor" json:"supervisor"`
}
