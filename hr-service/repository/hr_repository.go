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

func (r *HRRepository) Get(id int32) (*model.Employee, error) {
	row := r.DB.QueryRow("SELECT id, name, department, role, doj FROM employees WHERE id = $1", id)

	emp := &model.Employee{}
	err := row.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Role, &emp.DOJ)
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (r *HRRepository) List() ([]*model.Employee, error) {
	rows, err := r.DB.Query("SELECT id, name, department, role, doj FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		emp := &model.Employee{}
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Role, &emp.DOJ); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}
