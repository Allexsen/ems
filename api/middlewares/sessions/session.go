package session

import (
	"fmt"
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
			fmt.Println("No session")
			c.Redirect(http.StatusSeeOther, "/sign-in")
			return
		}

		fmt.Println("Hit c.Next()")
		c.Next()
	}
}
