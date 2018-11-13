package models

import (
	"time"
)

type Membership struct {
	BaseModel
	//Id          int    `gorm:"AUTO_INCREMENT"`
	Name        string `sql:"size:200"`
	User        User   `gorm:"foreignkey:UserRefer"`
	Expiry_date *time.Time
	Earn_Pts    string `sql:"TYPE:json"`
}
