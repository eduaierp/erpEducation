package main

import (
	"database/sql"
	"exam-service/handler"
	"exam-service/repository"
	"exam-service/server"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	repo := &repository.ExamRepository{DB: db}
	h := &handler.ExamHandler{Repo: repo}
	server.StartGRPCServer(h)
}
