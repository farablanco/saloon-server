// author : Kenneth Phang
package models

import "time"

type Hairdresser struct {
	Id         int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt  time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deletedAt"`
	Name       string     `gorm:"unique_index;not null" sql:"size:200" db:"name" json:"name"`
	YearsOfExp int64      `db:"years_of_exp" json:"yearsofEx"`
	Rating     int64      `db:"rating" json:"rating"`
}
