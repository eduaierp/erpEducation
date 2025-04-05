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

func (r *ExamRepository) GetByID(id int32) (*model.Exam, error) {
	row := r.DB.QueryRow("SELECT id, subject, date, type FROM exams WHERE id = $1", id)
	exam := &model.Exam{}
	err := row.Scan(&exam.ID, &exam.Subject, &exam.Date, &exam.Type)
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (r *ExamRepository) ListAll() ([]*model.Exam, error) {
	rows, err := r.DB.Query("SELECT id, subject, date, type FROM exams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exams []*model.Exam
	for rows.Next() {
		exam := &model.Exam{}
		err := rows.Scan(&exam.ID, &exam.Subject, &exam.Date, &exam.Type)
		if err != nil {
			return nil, err
		}
		exams = append(exams, exam)
	}
	return exams, nil
}
