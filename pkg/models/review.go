package models

import "time"

type Review struct {
	ID         int       `db:"review_id"`
	Employee   int       `db:"employee_id"`
	ReviewedBy int       `db:"reviewed_by_id"`
	Rating     int       `db:"rating_id"`
	Comment    string    `db:"comment"`
	Date       time.Time `db:"review_date"`
}
