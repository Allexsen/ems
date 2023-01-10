package referral

import (
	"fmt"
	"net/http"

	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func ValidateReferral(c *gin.Context) {
	fmt.Println("Referral Validation Invoked")
	referral := c.Param("referral")
	ok, err := CheckReferral(referral)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println("Reached c.File")
	c.File("../../../html/registration.html")
}

func NewReferral(c *gin.Context) {
	referral := GetNewReferral()

	db := database.GetDB()
	db.Exec("INSERT INTO referrals(code) VALUES(?)", referral)
	c.String(200, referral)
}
