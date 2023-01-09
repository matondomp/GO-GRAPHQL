package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (r *Course) Create(name string, description string, categoryId string) (Course, error) {
	id := uuid.New().String()
	_, err := r.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description}, nil
}

func (r *Course) FindAll() ([]Course, error) {

	rows, err := r.db.Query("SELECT id, name, description FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	courses := []Course{}

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description})
	}

	return courses, nil
}

func (r *Category) FindByCourseId(courseId string) (Category, error) {
	var id, name, description string
	err := r.db.QueryRow("SELECT c.id, c.name, c.description FROM Categories c JOIN courses cs ON c.id = cs.category_id WHERE cs.id = $id", courseId).Scan(&id, &name, &description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
