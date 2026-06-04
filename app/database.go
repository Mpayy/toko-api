package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Mpayy/toko-api/helper"
	"github.com/spf13/viper"
)

func NewDb(config *viper.Viper) *sql.DB {
	name := config.GetString("DATABASE_NAME")
	host := config.GetString("DATABASE_HOST")
	port := config.GetString("DATABASE_PORT")
	user := config.GetString("DATABASE_USER")
	pass := config.GetString("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
