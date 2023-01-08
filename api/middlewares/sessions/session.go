package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ph struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user ph
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("firstname", user.FirstName)
		session.Set("lastname", user.LastName)
		session.Save()

		c.Next()
	}
}
