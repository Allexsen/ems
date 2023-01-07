package main // Spin-off the app

import (
	"github.com/Allexsen/ems/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Initialize(router)
}
