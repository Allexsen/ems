package session

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func StoreSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		session.Set("email", c.PostForm("email"))
		session.Set("authenticated", "true")
		session.Save()
	}
}

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("email") == nil {
			c.File(os.Getenv("HTML_DIR") + "/index.html")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if !session.Get("authenticated").(bool) {
			c.File(os.Getenv("HTML_DIR") + "/index.html")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		session.Set("authenticated", "false")
		session.Save()
	}
}
