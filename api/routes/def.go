package routes

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	gin.ForceConsoleColor()
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_KEY")))
	r.Use(sessions.Sessions("sessions", store))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/profile/asvanidze12")
	})

	initAuth(r)
	initProfile(r)
	initReferral(r)

	r.Run()
}
