package referral

import (
	"fmt"
	"strconv"

	"github.com/Allexsen/ems/database"
)

func CheckReferral(refCode string) (bool, error) {
	fmt.Println("refCode=" + refCode)
	db := database.GetDB()
	row := db.QueryRow("SELECT is_used FROM referrals WHERE code=?", refCode)

	var isUsed bool
	err := row.Scan(&isUsed)
	fmt.Println("isUsed value=" + strconv.FormatBool(isUsed))
	if err != nil || isUsed {
		return false, err
	}

	return true, nil
}
