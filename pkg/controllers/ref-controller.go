package controllers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/Allexsen/ems/database"
	"github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
)

func getNewReferral() string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Doesn't even need to check if it's duplicate
	// chance of having duplicate is
	// (number of existing referrals)/62^62
	// Basically impossible
	var refCode [20]byte
	for K := 0; K < 20; K++ {
		refCode[K] = charSet[rand.Intn(len(charSet))]
	}

	return string(refCode[:])
}

func ValidateReferral() gin.HandlerFunc {
	return func(c *gin.Context) {
		referral := c.PostForm("referral")
		ok, err := models.CheckReferral(referral)
		if err != nil || !ok {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func NewReferral() gin.HandlerFunc {
	return func(c *gin.Context) {
		referral := getNewReferral()

		db := database.GetDB()
		db.Exec("INSERT INTO referrals(code) VALUES(?)", referral)
		c.String(200, referral)
	}
}
