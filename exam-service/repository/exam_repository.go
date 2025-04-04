package repository

import (
	"database/sql"
	"exam-service/model"
)

type ExamRepository struct {
	DB *sql.DB
}

func (r *ExamRepository) Create(exam *model.Exam) error {
	return r.DB.QueryRow(
		"INSERT INTO exams (subject, date, type) VALUES ($1, $2, $3) RETURNING id",
		exam.Subject, exam.Date, exam.Type,
	).Scan(&exam.ID)
}
