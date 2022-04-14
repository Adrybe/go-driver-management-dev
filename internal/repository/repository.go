package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewRepository() (*sql.DB, error) {
	connStr := "user=nugxkoeirydlkl password=b6bcdbd33f115a380de9e9c0f8e8bd9042320e5d85e16c975e6f5eafef52d41e host=ec2-3-217-113-25.compute-1.amazonaws.com dbname=d80vbmv0622q6b  sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error open connection: %s", err)
	}

	return db, err
}
