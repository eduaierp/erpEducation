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

func (r *AdmissionRepository) GetByID(id int32) (*model.Admission, error) {
	admission := &model.Admission{}
	err := r.DB.QueryRow(
		"SELECT id, student_name, course, status, applied_date FROM admissions WHERE id = $1",
		id,
	).Scan(&admission.ID, &admission.StudentName, &admission.Course, &admission.Status, &admission.AppliedDate)
	if err != nil {
		return nil, err
	}
	return admission, nil
}

func (r *AdmissionRepository) List() ([]*model.Admission, error) {
	rows, err := r.DB.Query("SELECT id, student_name, course, status, applied_date FROM admissions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admissions []*model.Admission
	for rows.Next() {
		a := &model.Admission{}
		err := rows.Scan(&a.ID, &a.StudentName, &a.Course, &a.Status, &a.AppliedDate)
		if err != nil {
			return nil, err
		}
		admissions = append(admissions, a)
	}
	return admissions, nil
}
