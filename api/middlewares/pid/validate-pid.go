package prof_id

import (
	"strconv"

	"github.com/Allexsen/ems/database"
	"github.com/Allexsen/ems/pkg/models"
)

func GetProfile(pid string) (*models.Employee, error) {
	db := database.GetDB()
	id, err := strconv.Atoi(pid)
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT first_name, last_name, email FROM employees WHERE employee_id=?", id)
	var employee models.Employee
	err = row.Scan(&employee.FirstName, &employee.LastName, &employee.Email)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}
