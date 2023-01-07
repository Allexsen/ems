package routes // profiles

import (
	"net/http"

	prof_id "github.com/Allexsen/ems/api/middlewares/pid"
	_ "github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
)

func initProfile(r *gin.Engine) {
	profile := r.Group("/profile")
	{
		profile.GET("", func(c *gin.Context) {
			c.String(200, "Self Profile")
		})

		profile.GET("/:pid", func(c *gin.Context) {
			employee, err := prof_id.GetProfile(c.Param("pid"))
			if err != nil {
				c.Abort()
				panic(err)
			}
			c.JSON(http.StatusOK, employee)
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
}
