package referral

import (
	"math/rand"
)

func GetNewReferral() string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Doesn't even need to check if it's duplicate
	// chance of having duplicate is
	// (number of existing referrals)/62^62
	// Basically impossible
	var refCode [20]byte
	for K := 0; K < 20; K++ {
		refCode[K] = charSet[rand.Intn(len(charSet))]
	}

	return string(refCode[:])
}
