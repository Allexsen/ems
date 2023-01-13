package session

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func StoreSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		session.Set("count", count)
		session.Set("email", c.PostForm("email"))
		session.Save()
	}
}

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("email") == nil {
			c.Redirect(http.StatusSeeOther, "/sign-in")
			c.Abort()
		}

		c.Next()
	}
}
