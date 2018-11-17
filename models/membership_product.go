package models

import "time"

type MembershipProduct struct {
	Id         int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt  time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deletedAt"`
	Product    Product    `gorm:"foreignkey:ProductRefer" db:"product_id" json:"product"`
	Membership Membership `gorm:"foreignkey:MembershipRefer" db:"membership_id" json:"membership"`
	Count      int64      `gorm:"unique_index;not null" db:"count" json:"count"`
}
