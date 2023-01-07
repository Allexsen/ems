package routes // define group routes

import (
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func Initialize(router *gin.Engine) {
	r = router
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/auth")
	})

	auth := r.Group("/auth")
	auth.Use(gin.Logger())
	{
		auth.GET("/", func(c *gin.Context) {
			c.String(200, "Successfully rerouted")
		})
	}

	router.Run()
}
