package models

import "time"

type UserPts struct {
	BaseModel
	User           User      `gorm:"foreignkey:UserRefer" db:"user_id" json:"user"`
	Pts            Points    `gorm:"foreignkey:PointsRefer" db:"pts_id" json:"points"`
	DateAcumulated time.Time `sql:"size:200" db:"date_accumulated" json:"dateAccumulated"`
	SubTotalPts    int64     `db:"sub_total_pts" json:"subTotalPts"`
}
