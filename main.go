package main

import (
	"database/sql"
	"github.com/elmanelman/sql-judge-api/store/postgresstore"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=judge password=judge dbname=judge sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}

	st := postgresstore.New(db)

	s := NewServer(st)
	s.Start()
}
