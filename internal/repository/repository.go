package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewRepository() (*sql.DB, error) {
	connStr := "user=rdqoeghpgmnegr password=aaa1b4f4ec51dab764cb5869ecf92ac79391a360712f3ecebec93400534b7289 host=ec2-52-73-155-171.compute-1.amazonaws.com dbname=d7tsslec2uqb97  sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error open connection: %s", err)
	}

	return db, err
}
