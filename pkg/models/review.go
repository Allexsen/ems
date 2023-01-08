package models

import "time"

type Review struct {
	ID         uint      `db:"review_id"`
	Employee   uint      `db:"employee_id"`
	ReviewedBy uint      `db:"reviewed_by_id"`
	Rating     uint      `db:"rating_id"`
	Comment    string    `db:"comment"`
	Date       time.Time `db:"review_date"`
}
