package models

import "time"

type Employee struct {
	ID          uint      `db:"employee_id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	MiddleName  string    `db:"middle_name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	PhoneNumber string    `db:"phone_number"`
	ManagerID   uint      `db:"manager_id"`
	HireDate    time.Time `db:"hire_date"`
	RetireDate  time.Time `db:"retirement_date"`
	EmpType     uint      `db:"employment_type"`
	PositionID  uint      `db:"position_id"`
	RefCode     [20]byte  `db:"referral_code"`
}
