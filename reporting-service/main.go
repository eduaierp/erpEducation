package main

import (
	"database/sql"
	"reporting-service/handler"
	"reporting-service/repository"
	"reporting-service/server"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.ReportRepository{DB: db}
	h := &handler.ReportHandler{Repo: repo}
	server.StartGRPCServer(h)
}
