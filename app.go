package main

import(
	"context"
	"database/sql"
	"log"
	"reflect"

	"tutorial.sqlc.dev/app/tutorial"
	_ "github.com/lib/pq"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=root password=root dbname=pqgotest sslmode=disable")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	//listar todos os autores
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	//criar um autor
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kerninghan",
		Bio : sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	//pegando o autor que acabamos de inserir
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}
	
	//prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}