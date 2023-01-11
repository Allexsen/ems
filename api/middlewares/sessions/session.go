package session

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		firstName := session.Get("firstname")
		lastName := session.Get("lastname")

		if userID == nil || firstName == nil || lastName == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/sign-in")
			return
		}

		c.Next()
	}
}
