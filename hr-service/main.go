package main

import (
	"database/sql"
	"hr-service/handler"
	"hr-service/repository"
	"hr-service/server"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.HRRepository{DB: db}
	h := &handler.HRHandler{Repo: repo}
	server.StartGRPCServer(h)
}
