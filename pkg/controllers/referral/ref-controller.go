package referral

import (
	"net/http"
	"os"
	"path"

	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func ValidateReferral(c *gin.Context) {
	referral := c.PostForm("referral")
	ok, err := checkReferral(referral)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.File(path.Join(os.Getenv("HTML_DIR"), "/registration.html"))
}

func NewReferral(c *gin.Context) {
	referral := GetNewReferral()

	db := database.GetDB()
	db.Exec("INSERT INTO referrals(code) VALUES(?)", referral)
	c.String(200, referral)
}
