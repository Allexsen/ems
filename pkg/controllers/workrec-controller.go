package controllers

import (
	"net/http"

	"github.com/Allexsen/ems/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func GetWorkRecords(c *gin.Context) {
	wc, err := models.GetWorkRecords(c.PostForm("pid"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	wcJSON, err := json.Marshal(wc)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(200, wcJSON)
}
