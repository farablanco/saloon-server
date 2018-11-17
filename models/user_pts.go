package models

import "time"

type UserPts struct {
	Id             int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt      time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deletedAt"`
	User           User       `gorm:"foreignkey:UserRefer" db:"user_id" json:"user"`
	Pts            Points     `gorm:"foreignkey:PointsRefer" db:"pts_id" json:"points"`
	DateAcumulated time.Time  `sql:"size:200" db:"date_accumulated" json:"dateAccumulated"`
	SubTotalPts    int64      `db:"sub_total_pts" json:"subTotalPts"`
}
