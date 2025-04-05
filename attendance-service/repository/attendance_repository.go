package repository

import (
	"attendance-service/model"
	"database/sql"
)

type AttendanceRepository struct {
	DB *sql.DB
}

func (r *AttendanceRepository) Mark(attendance *model.Attendance) error {
	return r.DB.QueryRow(
		"INSERT INTO attendance (student_id, date, status) VALUES ($1, $2, $3) RETURNING id",
		attendance.StudentID, attendance.Date, attendance.Status,
	).Scan(&attendance.ID)
}

func (r *AttendanceRepository) GetByID(id int32) (*model.Attendance, error) {
	row := r.DB.QueryRow("SELECT id, student_id, date, status FROM attendance WHERE id = $1", id)
	var a model.Attendance
	err := row.Scan(&a.ID, &a.StudentID, &a.Date, &a.Status)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AttendanceRepository) ListByStudentID(studentID int32) ([]*model.Attendance, error) {
	rows, err := r.DB.Query("SELECT id, student_id, date, status FROM attendance WHERE student_id = $1", studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*model.Attendance
	for rows.Next() {
		var a model.Attendance
		if err := rows.Scan(&a.ID, &a.StudentID, &a.Date, &a.Status); err != nil {
			return nil, err
		}
		records = append(records, &a)
	}

	return records, nil
}
    