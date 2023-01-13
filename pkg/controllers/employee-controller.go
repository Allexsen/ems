package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		emp, err := models.GetEmployee(c.Param("pid"))
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		empJSON, err := json.Marshal(emp)
		if err != nil {
			fmt.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, string(empJSON))
	}
}

func NewEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp models.Employee
		emp.RefCode = c.PostForm("referral")
		emp.Email = c.PostForm("email")
		emp.Password = c.PostForm("password")
		emp.FirstName = c.PostForm("firstname")
		emp.MiddleName = c.PostForm("middlename")
		emp.LastName = c.PostForm("lastname")
		emp.PhoneNumber = c.PostForm("phone_number")
		emp.HireDate = time.Now().UTC()

		// THIS IS FOR TESTING PURPOSES, REMOVE LATER
		emp.EmpType = 1
		emp.PositionID = 1
		// UP TO HERE

		pswdHash, err := bcrypt.GenerateFromPassword([]byte(emp.Password), 10)
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
}
