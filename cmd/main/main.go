package main // Spin-off the app

import (
	"fmt"

	"github.com/Allexsen/ems/api/routes"
	"github.com/Allexsen/ems/database"
)

func main() {
	routes.InitDef()

	db := database.GetDB()
	defer func() {
		db.Close()
		fmt.Println("[MySQL] Database closed & disconnected successfully..")
	}()
}
