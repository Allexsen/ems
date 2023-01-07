package database // configure database connection to the system

import (
	connection "github.com/Allexsen/ems/database/connect"
)

func init() {
	connection.ConnectDB(500, 100, 100, 2)
	// db := connection.GetDB()
}
