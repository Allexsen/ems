package referral

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func ValidateReferral(c *gin.Context) {
	referral := c.PostForm("referral")
	ok, err := CheckReferral(referral)
	if err != nil || !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fullpath := os.Getenv("FP_EMS")
	fmt.Println(fullpath)
	fullpath = path.Join(fullpath, "html", "registration.html")
	fmt.Println(fullpath)
	c.File(fullpath)
}

func NewReferral(c *gin.Context) {
	referral := GetNewReferral()

	db := database.GetDB()
	db.Exec("INSERT INTO referrals(code) VALUES(?)", referral)
	c.String(200, referral)
}
