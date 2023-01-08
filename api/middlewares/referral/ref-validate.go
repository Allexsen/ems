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
		if err == nil && !isUsed {
			c.Request.Method = "GET"
			c.Redirect(http.StatusTemporaryRedirect, "/referral/register")
			return
		}

		c.Next()
	}
}
