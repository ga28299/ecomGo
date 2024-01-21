package db

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	pgCon *sqlx.DB
)

func SetDB(conString string) (*sqlx.DB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := sqlx.ConnectContext(ctx, "postgres", conString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pgCon = client

	if err = pgCon.PingContext(ctx); err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to PG DB")

	pgCon.SetMaxOpenConns(25)
	pgCon.SetMaxIdleConns(25)
	pgCon.SetConnMaxLifetime(5 * time.Minute)
	return pgCon, nil
}

func CloseDB() {
	if pgCon != nil {

		err := pgCon.Close()
		if err != nil {
			log.Println("Error closing PostgreSQL database connection:", err)
		} else {
			log.Println("PostgreSQL database connection closed")
		}
	}
}
