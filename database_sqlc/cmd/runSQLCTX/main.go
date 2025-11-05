package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alexduzi/database_sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)

	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	category := CategoryParams{
		ID:          uuid.NewString(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend courses", Valid: true},
	}
	course := CourseParams{
		ID:          uuid.NewString(),
		Name:        "Golang course",
		Description: sql.NullString{String: "Learn golang", Valid: true},
		Price:       100.00,
	}

	courseDb := NewCourseDB(dbConn)

	err = courseDb.CreateCourseAndCategory(ctx, category, course)
	if err != nil {
		panic(err)
	}

	courses, err := courseDb.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("ID: %s\n", course.ID)
		fmt.Printf("Name: %s\n", course.Name)
		fmt.Printf("Category Name: %s\n", course.CategoryName)
		fmt.Printf("Description: %s\n", course.Description.String)
		fmt.Printf("Price: %.2f\n", course.Price)
		fmt.Println()
	}
}
