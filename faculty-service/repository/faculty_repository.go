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

func (r *FacultyRepository) GetByID(id int32) (*model.Faculty, error) {
	row := r.DB.QueryRow("SELECT id, name, email, department, designation FROM faculties WHERE id = $1", id)

	var f model.Faculty
	err := row.Scan(&f.ID, &f.Name, &f.Email, &f.Department, &f.Designation)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *FacultyRepository) Update(f *model.Faculty) error {
	_, err := r.DB.Exec(
		"UPDATE faculties SET name = $1, email = $2, department = $3, designation = $4 WHERE id = $5",
		f.Name, f.Email, f.Department, f.Designation, f.ID,
	)
	return err
}

func (r *FacultyRepository) Delete(id int32) error {
	_, err := r.DB.Exec("DELETE FROM faculties WHERE id = $1", id)
	return err
}

func (r *FacultyRepository) List() ([]*model.Faculty, error) {
	rows, err := r.DB.Query("SELECT id, name, email, department, designation FROM faculties")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faculties []*model.Faculty
	for rows.Next() {
		var f model.Faculty
		if err := rows.Scan(&f.ID, &f.Name, &f.Email, &f.Department, &f.Designation); err != nil {
			return nil, err
		}
		faculties = append(faculties, &f)
	}

	return faculties, nil
}
