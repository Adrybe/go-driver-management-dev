package repository

import (
	"database/sql"
	"log"
)

func NewRepository() (*sql.DB, error) {
	connStr := "user=rdqoeghpgmnegr dbname=d7tsslec2uqb97 host=ec2-52-73-155-171.compute-1.amazonaws.com password=aaa1b4f4ec51dab764cb5869ecf92ac79391a360712f3ecebec93400534b7289 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db, err
}
