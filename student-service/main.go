package main

import (
	"database/sql"
	"student-service/handler"
	"student-service/repository"
	"student-service/server"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.StudentRepository{DB: db}
	h := &handler.StudentHandler{Repo: repo}
	server.StartGRPCServer(h)
}
