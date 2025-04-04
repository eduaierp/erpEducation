package repository

import (
    "database/sql"
    "academic-service/model"
)

type AcademicRepository struct {
    DB *sql.DB
}

func (r *AcademicRepository) CreateProgram(program *model.Program) error {
    _, err := r.DB.Exec("INSERT INTO programs (name, description) VALUES ($1, $2)", program.Name, program.Description)
    return err
}
