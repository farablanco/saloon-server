package models

import (
	"time"

	"github.com/kenken64/saloon-server/utils"
)

type User struct {
	Id        int        `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	Firstname string     `sql:"size:100" db:"firstname" json:"firstName"`
	Lastname  string     `sql:"size:100" db:"lastname" json:"lastName"`
	Password  string     `sql:"size:200" db:"password" json:"password"`
	Email     string     `gorm:"unique_index;not null" db:"email" json:"email"`
	IsAdmin   bool       `sql:"size:1" db:"isAdmin" json:"isAdmin"`
}

func (u *User) HashPassword(plain string) (string, error) {
	return utils.HashAndSalt([]byte(plain))
}

func (u *User) CheckPassword(plain string) bool {
	err := utils.VerifyPassword(plain, []byte(u.Password))
	return err == nil
}
