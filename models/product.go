// author : Kenneth Phang
package models

import "time"

type Product struct {
	Id        int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	Name      string     `gorm:"unique_index;not null" sql:"size:200" db:"name" json:"name"`
	Price     float32    `sql:"type:decimal(18,2)" db:"price" json:"price"`
	Pts       int64      `db:"pts" json:"pts"`
	Remarks   string     `sql:"size:200" db:"remarks" json:"remarks"`
	Gender    string     `sql:"size:1" db:"gender" json:"gender"`
	Type      string     `sql:"size:1" db:"type" json:"type"`
	PhotoURL  string     `db:"photourl" json:"photourl"`
	PhotoURL2 string     `db:"photourl2" json:"photourl2"`
}
