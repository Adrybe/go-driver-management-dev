package repository

import (
	"database/sql"
	"log"
)

func NewRepository() *sql.DB {
	connStr := "user=rdqoeghpgmnegr dbname=d7tsslec2uqb97 host=ec2-52-73-155-171.compute-1.amazonaws.com password=aaa1b4f4ec51dab764cb5869ecf92ac79391a360712f3ecebec93400534b7289 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)

	defer db.Close()

	return db
}

func createTables(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE [IF NOT EXISTS] administrators (id text, username text, adminpassword text, authorized text);
	CREATE TABLE [IF NOT EXISTS] drivers (id text, username text, adminpassword text, driving text);`)
	if err != nil {
		log.Fatal(err)
	}
}
