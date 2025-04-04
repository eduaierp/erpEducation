package repository

import (
    "database/sql"
    "student-service/model"
)

type StudentRepository struct {
    DB *sql.DB
}

func (r *StudentRepository) Create(student *model.Student) error {
    _, err := r.DB.Exec("INSERT INTO students (name, email, course) VALUES ($1, $2, $3)", student.Name, student.Email, student.Course)
    return err
}
