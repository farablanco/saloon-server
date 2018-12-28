// author : Kenneth Phang
package models

import "time"

type Membership struct {
	Id        int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	Name      string     `gorm:"unique_index;not null" sql:"size:200" db:"name" json:"name"`
	PtsGiven  int64      `db:"pts_given" json:"PtsGiven"`
}
