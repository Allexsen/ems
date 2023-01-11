package authr

import (
	"net/http"

	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUser(c *gin.Context) {
	email := c.PostForm("email")

	db := database.GetDB()
	row := db.QueryRow("SELECT password FROM employees WHERE email=?", email)

	var pswdHash string
	if err := row.Scan(&pswdHash); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	password := c.PostForm("password")
	if err := bcrypt.CompareHashAndPassword([]byte(pswdHash), []byte(password)); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Next()
}
