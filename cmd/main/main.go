package main // Spin-off the app

import (
	"github.com/Allexsen/ems/api/routes/def.go"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	def.Initialize(router)
}
