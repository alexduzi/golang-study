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
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) Create(name, description, category string) (*Course, error) {
	id := uuid.NewString()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, category)
	if err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  category,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return []Course{}, err
	}
	defer rows.Close()

	courses := make([]Course, 0)

	for rows.Next() {
		var id, name, description, categoryId string
		rows.Scan(&id, &name, &description, &categoryId)
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
		})
	}

	return courses, nil
}
