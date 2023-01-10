package Referral

import (
	"github.com/Allexsen/ems/database"
)

func GetReferral(refCode string) (bool, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT is_used FROM referrals WHERE code=?", refCode)

	var isUsed bool
	err := row.Scan(&isUsed)
	if err != nil || isUsed {
		return 0, err
	}

	return 1, nil
}
