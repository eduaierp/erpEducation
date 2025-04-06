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

func (r *ReportRepository) GetByID(id int32) (*model.Report, error) {
	row := r.DB.QueryRow("SELECT id, title, content, created_at FROM reports WHERE id = $1", id)

	var report model.Report
	err := row.Scan(&report.ID, &report.Title, &report.Content, &report.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportRepository) List() ([]*model.Report, error) {
	rows, err := r.DB.Query("SELECT id, title, content, created_at FROM reports")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*model.Report
	for rows.Next() {
		var report model.Report
		if err := rows.Scan(&report.ID, &report.Title, &report.Content, &report.CreatedAt); err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}

	return reports, nil
}
