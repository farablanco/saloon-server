package models

type Membership struct {
	BaseModel
	Name     string `sql:"size:200" db:"name" json:"name"`
	PtsGiven string `sql:"TYPE:json" db:"pts_given" json:"PtsGiven"`
}
