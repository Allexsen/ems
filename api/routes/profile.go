package routes // profiles

import (
	"github.com/gin-gonic/gin"
)

func initProfile() {
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

	r.Run()
}
