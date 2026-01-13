//CONFIG - SETUP DATABASE CONNECTION

package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
)

// dito ba talaga to
var DB *sql.DB

func ConnectDB() {

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = os.Getenv("DBNET")
	cfg.Addr = os.Getenv("DBADDRESS")
	cfg.DBName = os.Getenv("DBNAME")
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected!")
}
