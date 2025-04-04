package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "academic-service/repository"
    "academic-service/handler"
    "academic-service/server"
)

func main() {
    connStr := "user=postgres dbname=erp sslmode=disable password=postgres"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    repo := &repository.AcademicRepository{DB: db}
    h := &handler.AcademicHandler{Repo: repo}
    server.StartGRPCServer(h)
}
