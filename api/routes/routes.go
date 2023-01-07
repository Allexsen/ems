package routes // define group routes

import (
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func Initialize(router *gin.Engine) {
	r = router
	gin.ForceConsoleColor()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/auth")
	})

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

	profile := r.Group("/profile")
	{
		profile.GET("", func(c *gin.Context) {
			c.String(200, "Self Profile")
		})

		profile.GET("/:pid", func(c *gin.Context) {
			c.String(200, "Employee Profile")
		})

		profile.GET("/:pid/history", func(c *gin.Context) {
			c.String(200, "Employee History")
		})

		profile.GET("/:pid/reviews", func(c *gin.Context) {
			c.String(200, "Employee Reviews")
		})

		profile.POST("/:pid/reviews/add", func(c *gin.Context) {
			c.String(200, "Add Review")
		})
	}

	router.Run()
}
