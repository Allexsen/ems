package models

import "time"

type WorkRecord struct {
	ID          uint      `db:"record_id"`
	EmployeeID  uint      `db:"employee_id"`
	Position    string    `db:"position"`
	Description string    `db:"description"`
	From        time.Time `db:"start_date"`
	To          time.Time `db:"end_date"`
}
