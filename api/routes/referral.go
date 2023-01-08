package routes // referrals

import (
	"net/http"

	referralmw "github.com/Allexsen/ems/api/middlewares/referral"
	"github.com/Allexsen/ems/pkg/controllers/referral"
	"github.com/gin-gonic/gin"
)

func initReferral(r *gin.Engine) {
	ref := r.Group("/referral")
	{
		ref.GET("/validate", func(c *gin.Context) {
			c.String(200, "Input Your Referral Code")
		})

		ref.POST("/validate/:referral", referralmw.ValidateReferral(), func(c *gin.Context) {

			c.Redirect(http.StatusTemporaryRedirect, "/referral/register")
		})

		ref.GET("/invalid", func(c *gin.Context) {
			c.String(200, "Invalid Referral Code")
		})

		ref.GET("/new", func(c *gin.Context) {
			refCode := referral.NewReferral()
			c.JSON(200, refCode)
		})

		ref.GET("/terminate", func(c *gin.Context) {
			c.String(200, "Active Referrals List")
		})

		ref.DELETE("/terminate/:referral", func(c *gin.Context) {
			c.String(200, "Referral Terminated")
		})

		ref.GET("/register", func(c *gin.Context) {
			c.String(200, "Registration")
		})
	}
}
