package main

import (
	"academic-service/handler"
	"academic-service/repository"
	"academic-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.AcademicRepository{DB: db}
	h := &handler.AcademicHandler{Repo: repo}
	server.StartGRPCServer(h)
}
