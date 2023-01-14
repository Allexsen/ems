package models

import "time"

type WorkRecord struct {
	ID          int       `db:"record_id"`
	EmployeeID  int       `db:"employee_id"`
	Position    string    `db:"position"`
	Description string    `db:"description"`
	From        time.Time `db:"start_date"`
	To          time.Time `db:"end_date"`
}
