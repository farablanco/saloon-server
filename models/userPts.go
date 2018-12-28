// author : Kenneth Phang
package models

import "time"

type UserPts struct {
	Id           int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt    time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deletedAt"`
	UserID       int64      `gorm:"foreignkey:UserID;association_foreignkey:Refer" db:"user_id" json:"userId"`
	ProductID    int64      `db:"product_id" json:"productId"`
	AllocatedPts int64      `db:"allocated_pts" json:"allocatedPts"`
}
