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

func log(msg string) { // Log DB msgs
	fmt.Println("[DB]:" + msg + "..")
}

func ConnectDB(mc uint, mo uint, mi uint, mlt uint) { // Connect to Database, set DB params
	ConnParams.MaxConns = mc
	ConnParams.MaxOpen = mo
	ConnParams.MaxIdle = mi
	ConnParams.MaxLifeTimeMins = mlt

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

	db = d
	// Set & Log DB settings
	db.SetConnMaxLifetime(time.Minute * time.Duration(ConnParams.MaxLifeTimeMins))
	SetMaxOpenConns(ConnParams.MaxOpen)
	SetMaxIdleConns(ConnParams.MaxIdle)
	LogDBConnSettings()
}

func GetDB() *sql.DB { // return DB instance
	return db
}

func SetMaxOpenConns(maxOpenConns uint) { // Set maximum allowed open connections
	if maxOpenConns > ConnParams.MaxConns { // Check if requested amount > max conns allowed, limit if so
		log("MaxConn: Invalid amount - limiting to" + strconv.Itoa(int(ConnParams.MaxConns)))
		maxOpenConns = ConnParams.MaxConns
	}

	ConnParams.MaxOpen = maxOpenConns
	MO := int(maxOpenConns)
	log("Changing allowed MaxOpenConns to" + strconv.Itoa(MO))
	db.SetMaxOpenConns(MO)

	if ConnParams.MaxOpen < ConnParams.MaxIdle { // Check if MaxIdle > MaxOpen, limit if so
		SetMaxIdleConns(uint(MO))
	}
}

func SetMaxIdleConns(maxIdleConns uint) { // Set maximum allowed idle connections
	var MI int
	if maxIdleConns > ConnParams.MaxOpen { // Check if requested amount > max open allowed, limit if so
		maxIdleConns = ConnParams.MaxOpen
		MI := int(ConnParams.MaxIdle)
		log("MaxIdle: Invalid amount - limiting to" + strconv.Itoa(MI))
	}

	ConnParams.MaxIdle = maxIdleConns
	MI = int(maxIdleConns)
	log("Changing allowed MaxIdleConns to" + strconv.Itoa(MI))
	db.SetMaxIdleConns(MI)
}

func LogDBConnSettings() { // Log current settings of db
	log("MaxOpenConns:" + strconv.Itoa(int(ConnParams.MaxOpen)))
	log("MaxIdleConns:" + strconv.Itoa(int(ConnParams.MaxIdle)))
	log("MaxConnLifetimeMins:" + strconv.Itoa(int(ConnParams.MaxLifeTimeMins)))
}
