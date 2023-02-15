package routes

import (
	"github.com/Allexsen/ems/pkg/controllers"
	_ "github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
)

func initProfile() {
	profile := r.Group("/profile")
	// This needs to be fixed
	// profile.Use(session.CheckSession())
	{
		profile.GET("/", func(c *gin.Context) {
			c.File("../../html/profile.html") // Should be replaced with id of the user and redirected to "/:pid"
		})

		profile.GET("/:pid", controllers.GetEmployee)

		profile.GET("/:pid/history", controllers.GetWorkRecords)

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
