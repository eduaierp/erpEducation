package main

import (
	"attendance-service/handler"
	"attendance-service/repository"
	"attendance-service/server"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Database not reachable: %v", err)
	}

	repo := &repository.AttendanceRepository{DB: db}
	h := &handler.AttendanceHandler{Repo: repo}
	server.StartGRPCServer(h)
}
