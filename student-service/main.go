package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "student-service/repository"
    "student-service/handler"
    "student-service/server"
)

func main() {
    connStr := "user=postgres dbname=erp sslmode=disable password=postgres"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    repo := &repository.StudentRepository{DB: db}
    h := &handler.StudentHandler{Repo: repo}
    server.StartGRPCServer(h)
}
