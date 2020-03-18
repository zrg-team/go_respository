package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// PORT server port
var (
	PORT      = 0
	SECRETKEY []byte
	DBDRIVER  = ""
	DBURL     = ""
	DBDEBUG   = false
	DBSYNC    = false
)

// Load the server PORT
func Load() {
	var err error
	var errDBDEBUG error
	var errDBSYNC error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	DBDEBUG, errDBDEBUG = strconv.ParseBool(os.Getenv("DB_DEBUG"))
	if errDBDEBUG != nil {
		log.Fatal(errDBDEBUG)
	}
	DBSYNC, errDBSYNC = strconv.ParseBool(os.Getenv("DB_SYNC"))
	if errDBSYNC != nil {
		log.Fatal(errDBSYNC)
	}

	SECRETKEY = []byte(os.Getenv("API_SECRET"))
}
