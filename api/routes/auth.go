package routes

import (
	"net/http"
	"os"

	authr "github.com/Allexsen/ems/api/middlewares/auth"
	session "github.com/Allexsen/ems/api/middlewares/session"
	"github.com/Allexsen/ems/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func initAuth() {
	auth := r.Group("/")
	{
		auth.GET("/sign-in", func(c *gin.Context) {
			c.File("../../html/index.html")
		})

		auth.POST("/sign-in", authr.CheckUser(), session.StoreSession(), func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/")
		})

		auth.GET("/sign-up", func(c *gin.Context) {
			c.File("../../html/sign-up.html")
		})

		auth.POST("/sign-up", controllers.ValidateReferral, controllers.NewEmployee, session.StoreSession(), func(c *gin.Context) {
			c.Redirect(http.StatusSeeOther, "/profile/")
		})

		auth.GET("/logout", session.Logout(), func(c *gin.Context) {
			c.File(os.Getenv("HTML_DIR") + "/index.html")
		})
	}
}
