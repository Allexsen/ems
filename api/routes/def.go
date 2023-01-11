package routes // define group routes

import (
	"net/http"

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
	store := cookie.NewStore([]byte("my-secret-key"))
	r.Use(sessions.Sessions("my-session", store))

	initAuth(r)
	initProfile(r)
	initReferral(r)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/profile/pid")
	})

	r.Run()
}
