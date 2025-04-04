package main

import (
	"admission-service/handler"
	"admission-service/repository"
	"admission-service/server"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=erp sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

	repo := &repository.AdmissionRepository{DB: db}
	h := &handler.AdmissionHandler{Repo: repo}
	server.StartGRPCServer(h)
}
