package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "finance-service/repository"
    "finance-service/handler"
    "finance-service/server"
)

func main() {
    connStr := "user=postgres dbname=erp sslmode=disable password=postgres"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    repo := &repository.FinanceRepository{DB: db}
    h := &handler.FinanceHandler{Repo: repo}
    server.StartGRPCServer(h)
}
