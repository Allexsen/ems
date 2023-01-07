package database // configure database connection to the system

import (
	"database/sql"

	connection "github.com/Allexsen/ems/database/connect"
)

func init() {
	connection.ConnectDB(500, 100, 100, 2)
}

func GetDB() *sql.DB {
	return connection.GetDB()
}
