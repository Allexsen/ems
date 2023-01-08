package routes // define group routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func Initialize(router *gin.Engine) {
	r = router
	initReferral(r)
	initProfile(r)
	gin.ForceConsoleColor()

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user_id", "user")
		session.Set("firstname", "name")
		session.Set("lastname", "surname")
		session.Save()
	})

	r.Run()
}
