package repository

import (
	"admission-service/model"
	"database/sql"
)

type AdmissionRepository struct {
	DB *sql.DB
}

func (r *AdmissionRepository) Apply(a *model.Admission) error {
	a.Status = "Pending"
	return r.DB.QueryRow(
		"INSERT INTO admissions (student_name, course, status, applied_date) VALUES ($1, $2, $3, $4) RETURNING id",
		a.StudentName, a.Course, a.Status, a.AppliedDate,
	).Scan(&a.ID)
}
