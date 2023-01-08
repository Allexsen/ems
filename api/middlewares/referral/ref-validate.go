package referralmw

import (
	"net/http"

	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func ValidateReferral() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		referral := c.Param("referral")
		q := "SELECT used FROM referrals WHERE code=" + referral
		row := db.QueryRow(q)

		var isUsed bool
		err := row.Scan(&isUsed)
		if err != nil || isUsed {
			c.String(http.StatusUnauthorized, "Invalid referral")
		}

		c.Next()
	}
}
