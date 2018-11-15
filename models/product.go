package models

type Product struct {
	BaseModel
	Name   string  `sql:"size:200" db:"name" json:"name"`
	price  float32 `sql:"type:decimal(18,2)" db:"price" json:"price"`
	status string  `db:"status" json:"status"`
}
