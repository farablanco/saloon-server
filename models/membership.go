package models

import (
	"time"
)

type Membership struct {
	BaseModel
	Name        string     `sql:"size:200" db:"name" json:"name"`
	User        User       `gorm:"foreignkey:UserRefer" db:"user_id" json:"user"`
	Expiry_date *time.Time `db:"expiry_date" json:"expiry_date"`
	Earn_Pts    string     `sql:"TYPE:json" db:"earn_pts" json:"earn_pts"`
}
