package main

import (
	"database/sql"
	"reporting-service/handler"
	"reporting-service/repository"
	"reporting-service/server"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=erp sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

	repo := &repository.ReportRepository{DB: db}
	h := &handler.ReportHandler{Repo: repo}
	server.StartGRPCServer(h)
}
