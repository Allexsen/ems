package referral

import (
	"math/rand"

	"github.com/Allexsen/ems/database"
)

func NewReferral() string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Doesn't even need to check if it's duplicate
	// chance of having duplicate is
	// (number of existing referrals)/62^62
	// Basically impossible
	var refCode [20]byte
	for K := 0; K < 20; K++ {
		refCode[K] = charSet[rand.Intn(len(charSet))]
	}

	db := database.GetDB()
	db.Query("INSERT INTO referrals(code) VALUES(?)", refCode)

	return string(refCode[:])
}
