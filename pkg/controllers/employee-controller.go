package controllers

import (
	"net/http"
	"time"

	"github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewEmployee(c *gin.Context) {
	var emp models.Employee
	emp.FirstName = c.PostForm("firstname")
	emp.LastName = c.PostForm("lastname")
	emp.MiddleName = c.PostForm("middlename")
	emp.Email = c.PostForm("email")
	emp.PhoneNumber = c.PostForm("phone_number")
	emp.HireDate = time.Now()
	emp.RefCode = c.PostForm("referral")

	pswdHash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), 10)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	emp.Password = string(pswdHash)
	err = models.NewEmployee(emp)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/profile/pid")
}
