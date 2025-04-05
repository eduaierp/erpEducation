package repository

import (
	"academic-service/model"
	"database/sql"
)

type AcademicRepository struct {
	DB *sql.DB
}

// Program CRUD
func (r *AcademicRepository) CreateProgram(program *model.Program) error {
	return r.DB.QueryRow("INSERT INTO programs (name, description) VALUES ($1, $2) RETURNING id", program.Name, program.Description).Scan(&program.ID)
}

func (r *AcademicRepository) GetProgram(id int32) (*model.Program, error) {
	program := &model.Program{}
	err := r.DB.QueryRow("SELECT id, name, description FROM programs WHERE id = $1", id).Scan(&program.ID, &program.Name, &program.Description)
	return program, err
}

func (r *AcademicRepository) UpdateProgram(program *model.Program) error {
	_, err := r.DB.Exec("UPDATE programs SET name=$1, description=$2 WHERE id=$3", program.Name, program.Description, program.ID)
	return err
}

func (r *AcademicRepository) DeleteProgram(id int32) error {
	_, err := r.DB.Exec("DELETE FROM programs WHERE id=$1", id)
	return err
}

func (r *AcademicRepository) ListPrograms() ([]*model.Program, error) {
	rows, err := r.DB.Query("SELECT id, name, description FROM programs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []*model.Program
	for rows.Next() {
		p := &model.Program{}
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return nil, err
		}
		programs = append(programs, p)
	}
	return programs, nil
}

// Course CRUD
func (r *AcademicRepository) CreateCourse(course *model.Course) error {
	return r.DB.QueryRow("INSERT INTO courses (title, code, program_id) VALUES ($1, $2, $3) RETURNING id", course.Title, course.Code, course.ProgramID).Scan(&course.ID)
}

func (r *AcademicRepository) GetCourse(id int32) (*model.Course, error) {
	course := &model.Course{}
	err := r.DB.QueryRow("SELECT id, title, code, program_id FROM courses WHERE id = $1", id).
		Scan(&course.ID, &course.Title, &course.Code, &course.ProgramID)
	return course, err
}

func (r *AcademicRepository) UpdateCourse(course *model.Course) error {
	_, err := r.DB.Exec("UPDATE courses SET title=$1, code=$2, program_id=$3 WHERE id=$4", course.Title, course.Code, course.ProgramID, course.ID)
	return err
}

func (r *AcademicRepository) DeleteCourse(id int32) error {
	_, err := r.DB.Exec("DELETE FROM courses WHERE id=$1", id)
	return err
}

func (r *AcademicRepository) ListCourses() ([]*model.Course, error) {
	rows, err := r.DB.Query("SELECT id, title, code, program_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		c := &model.Course{}
		err := rows.Scan(&c.ID, &c.Title, &c.Code, &c.ProgramID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}
