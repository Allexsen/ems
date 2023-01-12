package authr

import (
	"fmt"
	"net/http"

	"github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.PostForm("password")
		pswdHash, err := models.AuthEmployee(c.PostForm("email"))
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(pswdHash), []byte(password)); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
