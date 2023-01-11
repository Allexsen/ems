package referral

import (
	"github.com/Allexsen/ems/database"
)

func checkReferral(refCode string) (bool, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT is_used FROM referrals WHERE ref_code=?", refCode)

	var isUsed bool
	err := row.Scan(&isUsed)

	if err != nil || isUsed {
		return false, err
	}

	return true, nil
}
