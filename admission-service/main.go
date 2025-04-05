package main

import (
	"admission-service/handler"
	"admission-service/repository"
	"admission-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	repo := &repository.AdmissionRepository{DB: db}
	h := &handler.AdmissionHandler{Repo: repo}
	server.StartGRPCServer(h)
}
