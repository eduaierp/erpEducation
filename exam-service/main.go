package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "exam-service/repository"
    "exam-service/handler"
    "exam-service/server"
)

func main() {
    db, err := sql.Open("postgres", "user=postgres dbname=erp sslmode=disable password=postgres")
    if err != nil {
        panic(err)
    }

    repo := &repository.ExamRepository{DB: db}
    h := &handler.ExamHandler{Repo: repo}
    server.StartGRPCServer(h)
}
