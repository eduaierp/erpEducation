package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "hr-service/repository"
    "hr-service/handler"
    "hr-service/server"
)

func main() {
    db, err := sql.Open("postgres", "user=postgres dbname=erp sslmode=disable password=postgres")
    if err != nil {
        panic(err)
    }

    repo := &repository.HRRepository{DB: db}
    h := &handler.HRHandler{Repo: repo}
    server.StartGRPCServer(h)
}
