package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	pgCon *sqlx.DB
	esCon *elasticsearch.Client
)

func SetDB(pgConString string, esConString string) (*sqlx.DB, *elasticsearch.Client, error) {
	pgDB, err := setupPostgres(pgConString)
	if err != nil {
		return nil, nil, err
	}

	esClient, err := setupElasticsearch(esConString)
	if err != nil {
		return nil, nil, err
	}

	pgCon = pgDB
	esCon = esClient

	return pgCon, esCon, nil
}

func setupPostgres(conString string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := sqlx.ConnectContext(ctx, "postgres", conString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = client.PingContext(ctx); err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("Connected to PG DB ")

	client.SetMaxOpenConns(25)
	client.SetMaxIdleConns(25)
	client.SetConnMaxLifetime(5 * time.Minute)

	return client, nil
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

func setupElasticsearch(esConString string) (*elasticsearch.Client, error) {

	cfg := elasticsearch.Config{
		Addresses: []string{
			esConString,
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil, err
	}

	esCon = esClient
	res, err := esCon.Cluster.Health()
	if err != nil {
		log.Println("Elasticsearch connection error:", err)
		return nil, err
	}

	if res.IsError() {
		log.Printf("Elasticsearch cluster health request failed with status code: %d", res.StatusCode)
		return nil, fmt.Errorf("elasticsearch cluster health request failed with status code: %v", res.StatusCode)
	}

	log.Printf("Connected to Elasticsearch : %v", res)

	return esCon, nil
}
