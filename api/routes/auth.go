package routes

import "github.com/gin-gonic/gin"

func initAuth(r *gin.Engine) {
	auth := r.Group("/")
	auth.GET("/sign-up", func(c *gin.Context) {
		// c.String(200, "Registration Page HTML")
		c.File("../../html/sign-up.html")
	})

	auth.POST("/sign-up", func(c *gin.Context) {
		c.String(200, "Registration Submit")
	})
}
