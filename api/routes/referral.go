package routes // referrals

import (
	"github.com/Allexsen/ems/pkg/controllers/referral"
	"github.com/gin-gonic/gin"
)

func initReferral(r *gin.Engine) {
	ref := r.Group("/referral")
	{
		ref.GET("/validate", func(c *gin.Context) {
			c.String(200, "Input Your Referral Code")
		})

		ref.GET("/invalid", func(c *gin.Context) {
			c.String(200, "Invalid Referral Code")
		})

		ref.GET("/new", referral.NewReferral)

		ref.GET("/terminate", func(c *gin.Context) {
			c.String(200, "Active Referrals List")
		})

		ref.DELETE("/terminate/:referral", func(c *gin.Context) {
			c.String(200, "Referral Terminated")
		})

		ref.GET("/:referral", referral.ValidateReferral)

		ref.GET("/register", func(c *gin.Context) {
			c.String(200, "Registration Page")
		})

		ref.POST("/register", func(c *gin.Context) {
			c.String(200, "Registration Submit")
		})
	}
}
