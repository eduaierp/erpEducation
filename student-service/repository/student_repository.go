package repository

import (
	"database/sql"
	"student-service/model"
)

type StudentRepository struct {
	DB *sql.DB
}

func (r *StudentRepository) Create(student *model.Student) error {
	return r.DB.QueryRow("INSERT INTO students (name, email, course) VALUES ($1, $2, $3) RETURNING id",
		student.Name, student.Email, student.Course).Scan(&student.ID)
}

func (r *StudentRepository) Get(id int32) (*model.Student, error) {
	student := &model.Student{}
	err := r.DB.QueryRow("SELECT id, name, email, course FROM students WHERE id = $1", id).
		Scan(&student.ID, &student.Name, &student.Email, &student.Course)
	return student, err
}

func (r *StudentRepository) Update(student *model.Student) error {
	_, err := r.DB.Exec("UPDATE students SET name = $1, email = $2, course = $3 WHERE id = $4",
		student.Name, student.Email, student.Course, student.ID)
	return err
}

func (r *StudentRepository) Delete(id int32) error {
	_, err := r.DB.Exec("DELETE FROM students WHERE id = $1", id)
	return err
}

func (r *StudentRepository) List() ([]*model.Student, error) {
	rows, err := r.DB.Query("SELECT id, name, email, course FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*model.Student
	for rows.Next() {
		student := &model.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Course)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
