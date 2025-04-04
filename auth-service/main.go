package main

import (
	"auth-service/handler"
	"auth-service/repository"
	"auth-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=erp sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

	repo := &repository.UserRepository{DB: db}
	handler := &handler.AuthHandler{Repo: repo}
	server.StartGRPCServer(handler)
}
