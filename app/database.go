package app

import (
	"database/sql"
	"time"

	"github.com/spf13/viper"
)

func NewDB() *sql.DB {

	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	var (
		driver = "mysql"
		dsn    = config.GetString("APP_DSN")
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
