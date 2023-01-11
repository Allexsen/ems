package referral

import (
	"github.com/Allexsen/ems/database"
)

func checkReferral(refCode string) (bool, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT is_used FROM referrals WHERE code=?", refCode)

	var employee uint
	err := row.Scan(&employee)
	if err != nil || employee != 0 {
		return false, err
	}

	return true, nil
}
