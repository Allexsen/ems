package referral

import (
	"github.com/Allexsen/ems/database"
)

func CheckReferral(refCode string) (bool, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT is_used FROM referrals WHERE code=?", refCode)

	var isUsed bool
	err := row.Scan(&isUsed)
	if err != nil || isUsed {
		return false, err
	}

	return true, nil
}