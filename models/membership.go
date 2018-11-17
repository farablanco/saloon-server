package models

import "time"

type Membership struct {
	Id        int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	Name      string     `sql:"size:200" db:"name" json:"name"`
	PtsGiven  string     `sql:"TYPE:json" db:"pts_given" json:"PtsGiven"`
}
