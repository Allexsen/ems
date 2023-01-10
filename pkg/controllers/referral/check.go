package referral

import (
	"fmt"
	"strconv"

	"github.com/Allexsen/ems/database"
)

func CheckReferral(refCode string) (bool, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT is_used FROM referrals WHERE code=?", refCode)

	var isUsed bool
	err := row.Scan(&isUsed)
	fmt.Println("isUsed value=" + strconv.FormatBool(isUsed))
	fmt.Println("refCode=" + refCode)
	if err != nil || isUsed {
		return false, err
	}

	return true, nil
}
