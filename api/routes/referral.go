package routes // referrals

import "github.com/gin-gonic/gin"

func initReferral(r *gin.Engine) {
	ref := r.Group("/referral")
	{
		ref.GET("/validate", func(c *gin.Context) {
			c.String(200, "Input Your Referral Code")
		})

		ref.POST("/validate/:referral", func(c *gin.Context) {
			c.String(200, "Referral Validation")
		})

		ref.GET("/invalid", func(c *gin.Context) {
			c.String(200, "Invalid Referral Code")
		})

		ref.GET("/new", func(c *gin.Context) {
			c.String(200, "Generate Referral Code")
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
