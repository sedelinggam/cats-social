package postgresqlpkg

import (
	"cats-social/config"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
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

	db, err := sqlx.Connect("pgx", config.FormatDSN())
	db.SetMaxOpenConns(60)
	db.SetConnMaxLifetime(60 * time.Millisecond)
	if err != nil {
		log.Println("m=GetPool,msg=connection has failed", err)
	}
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
