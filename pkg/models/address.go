package models

type Address struct {
	ID         uint   `db:"address_id"`
	EmployeeID uint   `db:"employee_id"`
	Type       uint   `db:"address_type"`
	Country    string `db:"country"`
	State      string `db:"state"`
	City       string `db:"city"`
	Street     string `db:"street"`
	ZipCode    string `db:"postal_code"`
}
