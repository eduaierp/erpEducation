package repository

import (
	"database/sql"
	"reporting-service/model"
)

type ReportRepository struct {
	DB *sql.DB
}

func (r *ReportRepository) Create(report *model.Report) error {
	return r.DB.QueryRow(
		"INSERT INTO reports (title, content, created_at) VALUES ($1, $2, $3) RETURNING id",
		report.Title, report.Content, report.CreatedAt,
	).Scan(&report.ID)
}
