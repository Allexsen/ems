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
	MaxConns        uint
	MaxOpen         uint
	MaxIdle         uint
	MaxLifeTimeMins uint
}

func init() {
	ConnParams.MaxConns = 500
	ConnParams.MaxOpen = 100
	ConnParams.MaxIdle = 100
	ConnParams.MaxLifeTimeMins = 2
}

func log(msg string) { // Log DB msgs
	fmt.Println("[DB]:" + msg + "..")
}

func ConnectDB() { // Connect to Database, set DB params
	user := os.Getenv("SQL_USER")
	pswd := os.Getenv("SQL_PSWD")
	msn := user + ":" + pswd + "@/sql_ems"

	d, err := sql.Open("mysql", msn)
	if err != nil {
		panic(err)
	}

	defer func() {
		db.Close()
		log("Disconnected successfully")
	}()

	log("Connection established successfully")
	LogDBConnSettings()

	db = d
	// Set DB settings
	db.SetConnMaxLifetime(time.Minute * time.Duration(ConnParams.MaxLifeTimeMins))
	setMaxConnParams(uint(ConnParams.MaxOpen))
}

func GetDB() *sql.DB { // return DB instance
	return db
}

func setMaxConnParams(maxOpenConns uint) { // Set maximum allowed open connections
	if maxOpenConns > ConnParams.MaxConns { // Check if requested amount > max conns allowed, limit if so
		log("MaxConn: Invalid amount - limiting to" + strconv.Itoa(int(ConnParams.MaxConns)))
		ConnParams.MaxOpen = maxOpenConns
	}

	MO := int(ConnParams.MaxOpen)
	log("Changing allowed MaxOpen connections to" + strconv.Itoa(MO))
	db.SetMaxOpenConns(MO)

	if ConnParams.MaxOpen < ConnParams.MaxIdle { // Check if MaxIdle > MaxOpen, limit if so
		setMaxIdleConns(uint(MO))
	}
}

func setMaxIdleConns(maxIdleConns uint) { // Set maximum allowed idle connections
	MI := int(ConnParams.MaxIdle)
	if maxIdleConns > ConnParams.MaxOpen { // Check if requested amount > max open allowed, limit if so
		ConnParams.MaxIdle = ConnParams.MaxOpen
		MI = int(ConnParams.MaxIdle)
		log("MaxIdle: Invalid amount - limiting to" + strconv.Itoa(MI))
	}

	log("Changing allowed MaxIdle connections to" + strconv.Itoa(MI))
	db.SetMaxIdleConns(MI)
}

func LogDBConnSettings() { // Log current settings of db
	log("Maximum open connections:" + strconv.Itoa(int(ConnParams.MaxOpen)))
	log("Maximum idle connections:" + strconv.Itoa(int(ConnParams.MaxIdle)))
	log("Maximum connection lifetime(minutes):" + strconv.Itoa(int(ConnParams.MaxLifeTimeMins)))
}
