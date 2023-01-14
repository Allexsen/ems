package session

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func StoreSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		session.Set("authenticated", "true")
		session.Save()
	}
}

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("authenticated") != "true" {
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