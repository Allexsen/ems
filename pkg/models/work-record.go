package models

import (
	"strconv"
	"time"

	"github.com/Allexsen/ems/database"
)

type WorkRecord struct {
	ID          int       `db:"record_id"`
	EmployeeID  int       `db:"employee_id"`
	Position    string    `db:"position"`
	Description string    `db:"description"`
	From        time.Time `db:"start_date"`
	To          time.Time `db:"end_date"`
}

func GetWorkRecords(pidstr string) (WorkRecord, error) {
	var wc WorkRecord
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		return wc, err
	}
	wc.EmployeeID = pid

	db := database.GetDB()
	q := `SELECT position, description, start_date, end_date FROM work_records WHERE employee_id=?`
	row := db.QueryRow(q, wc.EmployeeID)
	err = row.Scan(&wc.Position, &wc.Description, &wc.From, &wc.To)
	return wc, err
}
