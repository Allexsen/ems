package models

import "github.com/Allexsen/ems/database"

type Referral struct {
	Code     string `db:"ref_code" json:"referral"`
	IsUsed   bool   `db:"is_used" json:"is_used"`
	UsedByID int    `db:"used_by_id" json:"used_by_id"`
}

func CheckReferral(referral string) (bool, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT * FROM referrals WHERE referral=?", referral)
	err := row.Scan(&referral)
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewReferral(referral string) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO referrals(referral) VALUES(?)", referral)
	return err
}

func TerminateReferral(referral string) error {
	db := database.GetDB()
	_, err := db.Exec("REMOVE FROM referrals WHERE referral=?", referral)

	return err
}
