package routes // referrals

import (
	"github.com/Allexsen/ems/pkg/controllers/referral"
	"github.com/gin-gonic/gin"
)

func initReferral(r *gin.Engine) {
	ref := r.Group("/referral")
	{
		ref.GET("/", func(c *gin.Context) {
			c.String(200, "Referral Input HTML")
		})

		ref.GET("/:referral", referral.ValidateReferral)

		ref.GET("/invalid", func(c *gin.Context) {
			c.String(200, "Invalid Referral HTML")
		})

		ref.GET("/new", referral.NewReferral)

		ref.GET("/terminate", func(c *gin.Context) {
			c.String(200, "Active Referrals HTML")
		})

		ref.DELETE("/terminate/:referral", func(c *gin.Context) {
			c.String(200, "Referral Terminated")
		})

		ref.GET("/register", func(c *gin.Context) {
			c.String(200, "Registration Page HTML")
		})

		ref.POST("/register", func(c *gin.Context) {
			c.String(200, "Registration Submit")
		})
	}
}
