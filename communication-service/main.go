package main

import (
	"communication-service/handler"
	"communication-service/repository"
	"communication-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.MessageRepository{DB: db}
	h := &handler.MessageHandler{Repo: repo}
	server.StartGRPCServer(h)
}
