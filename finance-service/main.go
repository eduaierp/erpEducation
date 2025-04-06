package main

import (
	"database/sql"
	"finance-service/handler"
	"finance-service/repository"
	"finance-service/server"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.FinanceRepository{DB: db}
	h := &handler.FinanceHandler{Repo: repo}
	server.StartGRPCServer(h)
}
