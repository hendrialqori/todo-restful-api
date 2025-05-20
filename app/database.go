package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	var (
		driver = "mysql"
		dsn    = "root:root@(localhost:3306)/todo_api?parseTime=true"
	)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(2 * time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	return db
}
