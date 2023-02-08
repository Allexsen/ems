package routes

import (
	"github.com/Allexsen/ems/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func initReferral() {
	ref := r.Group("/referral")
	{
		ref.GET("/new", controllers.NewReferral())

		ref.GET("/terminate", func(c *gin.Context) {
			c.String(200, "Active Referrals HTML")
		})

		ref.DELETE("/terminate/:referral", func(c *gin.Context) {
			c.String(200, "Referral Terminated")
		})
	}
}
