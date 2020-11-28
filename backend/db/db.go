package db

import (
	"database/sql"
	"fmt"
	"log"

	"../config"
)

func getConnection() *sql.DB {
	DB := config.DB
	db, err := sql.Open(fmt.Sprintf("%s", DB.DBType), fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DB.User, DB.Pass, DB.Addr, DB.Port, DB.DB))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
