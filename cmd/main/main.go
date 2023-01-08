package main // Spin-off the app

import (
	"fmt"

	"github.com/Allexsen/ems/api/routes"
	"github.com/Allexsen/ems/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Initialize(router)

	store := cookie.NewStore([]byte("my-secret-key"))
	router.Use(sessions.Sessions("my-session", store))

	db := database.GetDB()
	defer func() {
		db.Close()
		fmt.Println("[MySQL] Database closed & disconnected successfully..")
	}()
}
