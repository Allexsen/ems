package referral

import (
	"net/http"

	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func ValidateReferral(c *gin.Context) {
	db := database.GetDB()
	referral := c.Param("referral")
	q := "SELECT is_used FROM referrals WHERE code=?"
	row := db.QueryRow(q, referral)

	var isUsed bool
	err := row.Scan(&isUsed)
	if err != nil || isUsed {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		c.String(200, "Referral: Valid")
	}
}

func NewReferral(c *gin.Context) {
	referral := GetNewReferral()

	db := database.GetDB()
	db.Exec("INSERT INTO referrals(code) VALUES(?)", referral)
	c.String(200, referral)
}
