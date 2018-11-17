package models

import "time"

type PaymentItems struct {
	Id        int64      `gorm:"primary_key" db:"id" json:"id"`
	CreatedAt time.Time  `gorm:"not null;" db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"not null;" db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	Payment   Payment    `gorm:"foreignkey:PaymentRefer" db:"payment_id" json:"payment_id"`
	SubTotal  float32    `sql:"type:decimal(18,2);" db:"sub_total" json:"sub_total"`
	Product   Product    `gorm:"foreignkey:ProductRefer" db:"product_id" json:"product_id"`
}
