package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
)

func getNewReferral() string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Doesn't even need to check if it's duplicate
	// chance of having duplicate is
	// (number of existing referrals)/62^62
	// Basically impossible
	rand.Seed(time.Now().Unix())
	var refCode [20]byte
	for K := 0; K < 20; K++ {
		refCode[K] = charSet[rand.Intn(len(charSet))]
	}

	return string(refCode[:])
}

func ValidateReferral(c *gin.Context) {
	referral := c.PostForm("referral")
	ok, err := models.CheckReferral(referral)
	if err != nil || !ok {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func NewReferral(c *gin.Context) {
	referral := getNewReferral()

	err := models.NewReferral(referral)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.String(200, referral)
}
