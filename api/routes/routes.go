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

	auth := r.Group("/auth")
	{
		auth.GET("", func(c *gin.Context) {
			c.String(200, "Login")
		})

		auth.GET("/validate", func(c *gin.Context) {
			c.String(200, "Validate referral")
		})

		auth.GET("/register", func(c *gin.Context) {
			c.String(200, "Register")
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
