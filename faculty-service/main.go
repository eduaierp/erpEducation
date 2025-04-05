package main

import (
	"database/sql"
	"faculty-service/handler"
	"faculty-service/repository"
	"faculty-service/server"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	repo := &repository.FacultyRepository{DB: db}
	h := &handler.FacultyHandler{Repo: repo}
	server.StartGRPCServer(h)
}
