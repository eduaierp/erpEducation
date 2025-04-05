package main

import (
	"auth-service/handler"
	"auth-service/repository"
	"auth-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.UserRepository{DB: db}
	h := &handler.AuthHandler{Repo: repo}

	server.StartGRPCServer(h)
}
