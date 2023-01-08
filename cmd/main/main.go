package main // Spin-off the app

import (
	"fmt"

	"github.com/Allexsen/ems/api/routes"
	"github.com/Allexsen/ems/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Initialize(router)

	db := database.GetDB()
	defer func() {
		db.Close()
		fmt.Println("[MySQL] Database closed & disconnected successfully..")
	}()
}
