package models

type Product struct {
	BaseModel
	Name    string  `sql:"size:200" db:"name" json:"name"`
	Price   float32 `sql:"type:decimal(18,2)" db:"price" json:"price"`
	Status  string  `db:"status" json:"status"`
	Remarks string  `sql:"size:200" db:"remarks" json:"remarks"`
	Gender  string  `sql:"size:1" db:"gender" json:"gender"`
}
