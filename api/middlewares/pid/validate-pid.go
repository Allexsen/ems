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

	row := db.QueryRow("SELECT * FROM employees WHERE employee_id = ?", uint(id))
	var employee models.Employee
	err = row.Scan(&employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}
