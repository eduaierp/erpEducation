package repository

import (
	"database/sql"
	"hr-service/model"
)

type HRRepository struct {
	DB *sql.DB
}

func (r *HRRepository) Add(emp *model.Employee) error {
	return r.DB.QueryRow(
		"INSERT INTO employees (name, department, role, doj) VALUES ($1, $2, $3, $4) RETURNING id",
		emp.Name, emp.Department, emp.Role, emp.DOJ,
	).Scan(&emp.ID)
}
