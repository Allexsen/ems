package routes // define group routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func Initialize(r *gin.Engine) {
	router = r
}

// func
