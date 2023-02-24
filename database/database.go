package database

import (
	"database/sql"
	"os"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PASSWORD")+"@tcp(localhost:"+os.Getenv("PORT")+")/"+os.Getenv("DATABASE"))
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CloseDB(db *sql.DB) {
	defer db.Close()
}
