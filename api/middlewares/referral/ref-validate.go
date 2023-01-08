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
		q := "SELECT is_used FROM referrals WHERE code=?"
		row := db.QueryRow(q, referral)

		var isUsed bool
		err := row.Scan(&isUsed)
		if err != nil || isUsed {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
