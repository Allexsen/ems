package models

import "github.com/Allexsen/ems/database"

type Referral struct {
	Code     string `db:"ref_code" json:"referral"`
	IsUsed   bool   `db:"is_used" json:"is_used"`
	UsedByID uint   `db:"used_by_id" json:"used_by_id"`
}

func CheckReferral(refCode string) (bool, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT is_used FROM referrals WHERE ref_code=?", refCode)

	var isUsed bool
	err := row.Scan(&isUsed)

	if err != nil || isUsed {
		return false, err
	}

	return true, nil
}

func NewReferral(referral string) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO referrals(code) VALUES(?)", referral)
	return err
}
