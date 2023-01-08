package referral

import "math/rand"

func NewReferral() string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	charSet := lower + upper + numbers

	var refCode [20]byte
	for K := 0; K < 20; K++ {
		refCode[K] = charSet[rand.Intn(len(charSet))]
	}

	return string(refCode[:])
}
