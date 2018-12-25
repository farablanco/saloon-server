// author : Kenneth Phang
package models

import (
	"errors"
	"time"

	"../utils"
)

type User struct {
	Id                     int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt              time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt              time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt              *time.Time `db:"deleted_at" json:"deletedAt"`
	Firstname              string     `sql:"size:100" db:"firstname" json:"firstName"`
	Lastname               string     `sql:"size:100" db:"lastname" json:"lastName"`
	Password               string     `sql:"size:200" db:"password" json:"password"`
	Email                  string     `gorm:"unique_index;not null" db:"email" json:"email"`
	Username               string     `gorm:"unique_index;not null" db:"username" json:"username"`
	IsAdmin                bool       `sql:"size:1" db:"isAdmin" json:"isAdmin"`
	PtsBalance             int64      `db:"pts_balace" json:"ptsBalance"`
	ExpiryDate             *time.Time `db:"expiry_date" json:"expiryDate"`
	ChangePasswordDate     *time.Time `db:"change_password_date" json:"changePasswordDate"`
	ResetPasswordDate      *time.Time `db:"reset_password_date" json:"resetPasswordDate"`
	ResetPasswordToken     string     `db:"reset_password_token" json:"resetPasswordToken"`
	Membership             int64      `db:"membership_id" json:"membership"`
	Outlet                 int64      `db:"outlet_id" json:"outlet"`
	CutCnt                 int64      `db:"c_cnt" json:"cutCnt"`
	TreatmentCount         int64      `db:"t_cnt" json:"treatment_cnt"`
	HairlossTreatmentCount int64      `db:"ht_cnt" json:"hairloss_treatment_cnt"`
	ContactNo              string     `gorm:"unique_index;not null" sql:"size:20" db:"contact_no" json:"contactNo"`
	Image                  string     `sql:"size:1000" db:"image" json:"image"`
	Bio                    string     `sql:"size:2000" db:"bio" json:"bio"`
}

func (u *User) HashPassword(plain string) error {
	if len(plain) == 0 {
		return errors.New("password should not be empty!")
	}
	passwordHash, _ := utils.HashAndSalt([]byte(plain))
	u.Password = string(passwordHash)
	return nil
}

func (u *User) CheckPassword(plain string) bool {
	err := utils.VerifyPassword(plain, []byte(u.Password))
	return err == nil
}
