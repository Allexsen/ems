package routes // define group routes

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func Initialize(router *gin.Engine) {
	r = router
	gin.ForceConsoleColor()
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_KEY")))
	r.Use(sessions.Sessions("sessions", store))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/profile/asvanidze12@gmail.com")
	})

	initAuth(r)
	initProfile(r)
	initReferral(r)

	r.Run()
}
