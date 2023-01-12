package routes

import (
	"net/http"

	authr "github.com/Allexsen/ems/api/middlewares/auth"
	"github.com/Allexsen/ems/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func initAuth(r *gin.Engine) {
	auth := r.Group("/")
	{
		auth.GET("/sign-in", func(c *gin.Context) {
			c.File("../../html/index.html")
		})

		auth.POST("/sign-in", authr.CheckUser(), func(c *gin.Context) {
			c.Redirect(http.StatusSeeOther, "/profile/:pid")
		})

		auth.GET("/sign-up", func(c *gin.Context) {
			c.File("../../html/sign-up.html")
		})

		auth.POST("/sign-up", controllers.ValidateReferral(), controllers.NewEmployee)
	}
}
