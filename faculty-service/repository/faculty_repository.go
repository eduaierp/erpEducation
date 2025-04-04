package repository

import (
	"database/sql"
	"faculty-service/model"
)

type FacultyRepository struct {
	DB *sql.DB
}

func (r *FacultyRepository) Create(f *model.Faculty) error {
	return r.DB.QueryRow(
		"INSERT INTO faculties (name, email, department, designation) VALUES ($1, $2, $3, $4) RETURNING id",
		f.Name, f.Email, f.Department, f.Designation,
	).Scan(&f.ID)
}
