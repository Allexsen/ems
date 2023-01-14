package models

import (
	"github.com/Allexsen/ems/database"
)

type Employee struct {
	ID          int    `db:"employee_id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	MiddleName  string `db:"middle_name" json:"middle_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	TeamID      int    `db:"team_id" json:"team_id"`
	EmpType     int    `db:"employment_type" json:"employment_type"`
	PositionID  int    `db:"position_id" json:"position_id"`
}

func GetEmployee(email string) (Employee, error) {
	db := database.GetDB()
	q := `SELECT first_name, middle_name, last_name, phone_number FROM employees WHERE email=?`
	row := db.QueryRow(q, email)
	var emp Employee
	err := row.Scan(&emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.PhoneNumber)

	return emp, err
}

func NewEmployee(emp Employee) error {
	db := database.GetDB()
	// q := `INSERT INTO
	// 	employees(first_name, middle_name, last_name, email, password, phone_number, hire_date, referral_code)
	// 	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	// _, err := db.Exec(q, emp.FirstName, emp.MiddleName, emp.LastName, emp.Email, emp.Password,
	// 	emp.PhoneNumber, emp.HireDate, emp.RefCode)

	q := `INSERT INTO 
		employees(first_name, middle_name, last_name, email, password, phone_number, hire_date, employment_type, position_id)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q, emp.FirstName, emp.MiddleName, emp.LastName, emp.Email, emp.Password,
		emp.PhoneNumber, emp.EmpType, emp.PositionID)

	return err
}

func AuthEmployee(email string) (string, error) {
	db := database.GetDB()
	q := `SELECT password FROM employees WHERE email=?`
	row := db.QueryRow(q, email)

	var pswdHash string
	err := row.Scan(&pswdHash)

	return pswdHash, err
}
