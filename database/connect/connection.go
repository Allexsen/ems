package connection // Establish connection to the database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db         *sql.DB
	ConnParams Conns
)

type Conns struct { // Connection params
	MaxConns        int
	MaxOpen         int
	MaxIdle         int
	MaxLifeTimeMins int
}

func log(msg string) { // Log DB msgs
	fmt.Println("[MySQL] " + msg)
}

func ConnectDB(mc int, mo int, mi int, mlt int) { // Connect to Database, set DB params
	ConnParams.MaxConns = mc
	ConnParams.MaxOpen = mo
	ConnParams.MaxIdle = mi
	ConnParams.MaxLifeTimeMins = mlt

	user := os.Getenv("SQL_USER")
	pswd := os.Getenv("SQL_PSWD")
	adrs := os.Getenv("SQL_ADRS")
	msn := user + ":" + pswd + "@tcp(" + adrs + ")/sql_ems"

	d, err := sql.Open("mysql", msn)
	if err != nil {
		panic(err)
	}

	db = d
	log("Connection established successfully at " + adrs)

	// Set & Log DB settings
	db.SetConnMaxLifetime(time.Minute * time.Duration(ConnParams.MaxLifeTimeMins))
	SetMaxOpenConns(ConnParams.MaxOpen)
	SetMaxIdleConns(ConnParams.MaxIdle)
	LogDBConnSettings()
}

func GetDB() *sql.DB { // return DB instance
	return db
}

func SetMaxOpenConns(maxOpenConns int) { // Set maximum allowed open connections
	if maxOpenConns > ConnParams.MaxConns { // Check if requested amount > max conns allowed, limit if so
		log("MaxConn: Invalid amount - limiting to " + strconv.Itoa(int(ConnParams.MaxConns)))
		maxOpenConns = ConnParams.MaxConns
	}

	ConnParams.MaxOpen = maxOpenConns
	log("Changing allowed MaxOpenConns to " + strconv.Itoa(maxOpenConns))
	db.SetMaxOpenConns(maxOpenConns)

	if ConnParams.MaxOpen < ConnParams.MaxIdle { // Check if MaxIdle > MaxOpen, limit if so
		SetMaxIdleConns(int(maxOpenConns))
	}
}

func SetMaxIdleConns(maxIdleConns int) { // Set maximum allowed idle connections
	if maxIdleConns > ConnParams.MaxOpen { // Check if requested amount > max open allowed, limit if so
		maxIdleConns = ConnParams.MaxOpen
		log("MaxIdle: Invalid amount - limiting to " + strconv.Itoa(ConnParams.MaxIdle))
	}

	ConnParams.MaxIdle = maxIdleConns
	log("Changing allowed MaxIdleConns to " + strconv.Itoa(maxIdleConns))
	db.SetMaxIdleConns(maxIdleConns)
}

func LogDBConnSettings() { // Log current settings of db
	log("Current connection settings:")
	log("MaxOpenConns: " + strconv.Itoa(ConnParams.MaxOpen))
	log("MaxIdleConns: " + strconv.Itoa(ConnParams.MaxIdle))
	log("MaxConnLifetimeMins: " + strconv.Itoa(ConnParams.MaxLifeTimeMins))
}
