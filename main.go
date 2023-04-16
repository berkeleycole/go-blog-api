package main

import (
	"blog-api/http"
	"blog-api/mongo"
	"blog-api/writings"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	db := mongo.InitDB(ctx)

	// init repo
	writings_repo := mongo.NewWritingsRepo(db)

	// TODO: write test env
	// writing_repo := mockdb.New()

	// init service
	writings_service := writings.New(&writings_repo)

	// init server
	server := http.New(writings_service)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
