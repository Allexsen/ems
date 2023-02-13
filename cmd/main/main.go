package main // Spin-off the app

import (
	"fmt"

	"github.com/Allexsen/ems/database"
)

func main() {
	db := database.GetDB()
	defer func() {
		db.Close()
		fmt.Println("[MySQL] Database closed & disconnected successfully..")
	}()
}
