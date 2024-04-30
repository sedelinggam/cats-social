package postgresqlpkg

import (
	"cats-social/config"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgreSQL() *sqlx.DB {
	config := config.PostgreSQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Sslmode:  os.Getenv("DB_PARAMS"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}

	db, err := sqlx.Open("postgres", config.FormatDSN())
	db.SetMaxIdleConns(18)
	db.SetConnMaxLifetime(2 * time.Minute)
	db.SetMaxOpenConns(18)
	if err != nil {
		log.Println("m=GetPool,msg=connection has failed", err)
	}
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
