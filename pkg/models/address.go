package models

type Address struct {
	ID         int    `db:"address_id"`
	EmployeeID int    `db:"employee_id"`
	Type       int    `db:"address_type"`
	Country    string `db:"country"`
	State      string `db:"state"`
	City       string `db:"city"`
	Street     string `db:"street"`
	ZipCode    string `db:"postal_code"`
}
