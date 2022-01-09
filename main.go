package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	db "github.com/GavukaAlexandr/ent-sandbox/db"

	_ "github.com/lib/pq"

	"go.uber.org/zap"
)

func run() error {
	ctx := context.Background()

	dataBase, err := sql.Open("postgres", "user=sandbox_user dbname=sandbox_db sslmode=disable password=1")
	if err != nil {
		return err
	}

	queries := db.New(dataBase)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	zap.L().Info("STARTED APP")
	if err := run(); err != nil {
		zap.S().Fatal(err)
	}
}
