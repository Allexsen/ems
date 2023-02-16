package routes

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	r     *gin.Engine
	store cookie.Store
)

func InitDef() {
	gin.ForceConsoleColor()
	r = gin.Default()
	r.SetTrustedProxies(nil)

	store = cookie.NewStore([]byte(os.Getenv("COOKIE_KEY")))
	r.Use(sessions.Sessions("sessions", store))

	initAuth()
	initProfile()
	initReferral()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/profile/7")
	})

	r.Run()
}
