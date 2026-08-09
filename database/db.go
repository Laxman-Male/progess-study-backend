package database

import (
	// "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect(user, pass, host, name string) *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, name)
	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("error in opening DB-", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("error connecting DB-", err)
	}
	fmt.Println("connected")

	return DB

}
