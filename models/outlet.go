package models

type Outlet struct {
	BaseModel
	Name       string `sql:"size:200" db:"name" json:"name"`
	PostalCode int    `db:"postal_code" json:"postal_code"`
	ContactNo  string `sql:"size:45" db:"contact_no" json:"contactNo"`
	Email      string `sql:"size:100" db:"email" json:"email"`
}
