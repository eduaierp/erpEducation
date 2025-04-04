package main

import (
	"communication-service/handler"
	"communication-service/repository"
	"communication-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=erp sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

	repo := &repository.MessageRepository{DB: db}
	h := &handler.MessageHandler{Repo: repo}
	server.StartGRPCServer(h)
}
