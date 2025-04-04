package repository

import (
    "database/sql"
    "attendance-service/model"
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
