package main

import (
	"attendance-service/handler"
	"attendance-service/repository"
	"attendance-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=erp sslmode=disable password=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.AttendanceRepository{DB: db}
	h := &handler.AttendanceHandler{Repo: repo}
	server.StartGRPCServer(h)
}
