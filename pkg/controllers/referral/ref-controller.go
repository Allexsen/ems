package referral

import (
	"net/http"

	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func ValidateReferral(c *gin.Context) {
	referral := c.Param("referral")
	ok, err := CheckReferral(referral)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.String(200, "Referral: Valid")
}

func NewReferral(c *gin.Context) {
	referral := GetNewReferral()

	db := database.GetDB()
	db.Exec("INSERT INTO referrals(code) VALUES(?)", referral)
	c.String(200, referral)
}
