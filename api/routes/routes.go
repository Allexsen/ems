package routes // define group routes

import (
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
		c.String(200, "Login")
	})

	r.Run()
}
