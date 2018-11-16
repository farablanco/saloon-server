package models

type MembershipProduct struct {
	BaseModel
	Product    Product    `gorm:"foreignkey:ProductRefer" db:"product_id" json:"product"`
	Membership Membership `gorm:"foreignkey:MembershipRefer" db:"membership_id" json:"membership"`
	Count      int64      `gorm:"unique_index;not null" db:"count" json:"count"`
}
