package models

import (
	"time"

	"../utils"
)

type User struct {
	BaseModel
	Firstname              string     `sql:"size:100" db:"firstname" json:"firstName"`
	Lastname               string     `sql:"size:100" db:"lastname" json:"lastName"`
	Password               string     `sql:"size:200" db:"password" json:"password"`
	Email                  string     `gorm:"unique_index;not null" db:"email" json:"email"`
	IsAdmin                bool       `sql:"size:1" db:"isAdmin" json:"isAdmin"`
	PtsBalance             int64      `db:"pts_balace" json:"ptsBalance"`
	ExpiryDate             *time.Time `db:"expiry_date" json:"expiry_date"`
	Membership             Membership `gorm:"foreignkey:MembershipRefer" db:"membership_id" json:"membership"`
	Outlet                 Outlet     `gorm:"foreignkey:OutletRefer" db:"outlet_id" json:"outlet"`
	CutCount               int64      `db:"cut_cnt" json:"cutCount"`
	TreatmentCount         int64      `db:"treatment_cnt" json:"treatmentCnt"`
	HairlossTreatmentCount int64      `db:"hairloss_treatment_cnt" json:"hairLossTreatmentCnt"`
	ContactNo              string     `sql:"size:20" db:"contact_no" json:"contactNo"`
}

func (u *User) HashPassword(plain string) (string, error) {
	return utils.HashAndSalt([]byte(plain))
}

func (u *User) CheckPassword(plain string) bool {
	err := utils.VerifyPassword(plain, []byte(u.Password))
	return err == nil
}
