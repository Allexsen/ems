package routes // profiles

import (
	session "github.com/Allexsen/ems/api/middlewares/sessions"
	_ "github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
)

func initProfile(r *gin.Engine) {
	profile := r.Group("/profile")
	profile.Use(session.CheckSession())
	{
		profile.GET("/:pid", func(c *gin.Context) {
			c.File("../../html/profile.html")
		})

		profile.GET("/:pid/history", func(c *gin.Context) {
			c.File("../../html/work-history.html")
		})

		profile.GET("/:pid/reviews", func(c *gin.Context) {
			c.File("../../html/reviews.html")
		})

		profile.GET("/:pid/reviews/add", func(c *gin.Context) {
			c.File("../../html/review-add.html")
		})

		profile.POST("/:pid/reviews/add", func(c *gin.Context) {
			c.File("../../html/reviews.html")
		})
	}
}
