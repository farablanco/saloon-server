package models

import "time"

type Points struct {
	Id        int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	Product   Product    `gorm:"foreignkey:ProductRefer" db:"product_id" json:"product_id"`
	Pts       int        `db:"pts" json:"pts"`
	Price     float32    `sql:"type:decimal(18,2);" db:"price" json:"price"`
}
