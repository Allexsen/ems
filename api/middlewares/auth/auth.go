package authr

import (
	"net/http"

	"github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.PostForm("password")
		pswdHash, err := models.AuthEmployee(c.PostForm("email"), password)
		if err != nil {
			c.Redirect(http.StatusBadRequest, "/sign-in")
			c.Abort()
		}

		if bcrypt.CompareHashAndPassword([]byte(pswdHash), []byte(password)) != nil {
			c.Redirect(http.StatusUnauthorized, "/sign-in")
			c.Abort()
		}

		c.Next()
	}
}
