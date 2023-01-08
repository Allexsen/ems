package referralmw

import (
	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func ValidateReferral(c *gin.Context) bool {
	db := database.GetDB()
	referral := c.Param("referral")
	q := "SELECT is_used FROM referrals WHERE code=?"
	row := db.QueryRow(q, referral)

	var isUsed bool
	err := row.Scan(&isUsed)
	if err != nil || isUsed {
		return false
	}

	return true
}
