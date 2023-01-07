package routes // define group routes

import (
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func Initialize(router *gin.Engine) {
	initReferral()
	initProfile()
	r = router
	gin.ForceConsoleColor()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Login")
	})

	r.Run()
}
