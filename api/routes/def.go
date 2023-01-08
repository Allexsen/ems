package routes // define group routes

import (
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

	initReferral(r)
	initProfile(r)

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user_id", 2)
		session.Set("firstname", "name")
		session.Set("lastname", "surname")
		session.Save()
	})

	r.Run()
}
