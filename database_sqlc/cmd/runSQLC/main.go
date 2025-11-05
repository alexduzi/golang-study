package main

import (
	"context"
	"database/sql"

	"github.com/alexduzi/database_sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend 1",
	// 	Description: sql.NullString{String: "Studying backend 1", Valid: true},
	// })
	// if err != nil {
	// 	fmt.Printf("%#v", err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, cat := range categories {
	// 	fmt.Printf("ID: %s\n", cat.ID)
	// 	fmt.Printf("Name: %s\n", cat.Name)
	// 	fmt.Printf("Description: %s\n", cat.Description.String)
	// 	fmt.Println()
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "feaccdbc-1e2b-426c-8d7d-9f5bce55772a",
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, "feaccdbc-1e2b-426c-8d7d-9f5bce55772a")
	if err != nil {
		panic(err)
	}

	queries.CreateCourse(ctx, db.CreateCourseParams{
		ID: uuid.New().String(),
	})
}
