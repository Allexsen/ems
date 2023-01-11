package models

import (
	"time"

	"github.com/Allexsen/ems/database"
)

type Employee struct {
	ID          uint      `db:"employee_id" json:"id"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	MiddleName  string    `db:"middle_name" json:"middle_name"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"password"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	ManagerID   uint      `db:"manager_id" json:"manager_id"`
	HireDate    time.Time `db:"hire_date" json:"hire_date"`
	RetireDate  time.Time `db:"retirement_date" json:"retirement_date"`
	EmpType     uint      `db:"employment_type" json:"employment_type"`
	PositionID  uint      `db:"position_id" json:"position_id"`
	RefCode     string    `db:"referral_code" json:"referral_code"`
}

func NewEmployee(employee Employee) error {
	db := database.GetDB()
	q := `INSERT INTO 
			employees(first_name, last_name, middle_name, email, password, phone_number, hire_date, referral_code)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q, employee.FirstName, employee.LastName, employee.MiddleName, employee.Email, employee.Password,
		employee.PhoneNumber, employee.HireDate, employee.RefCode)

	return err
}
