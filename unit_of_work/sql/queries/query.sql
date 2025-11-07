-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = ?;

-- name: CreateCategory :execresult
INSERT INTO categories (name)
VALUES (?);

-- name: CreateCourse :exec
INSERT INTO courses (name, description, category_id)
VALUES (?, ?, ?);

-- name: ListCourses :many
SELECT c.*, ca.name as category_name
FROM courses c
JOIN categories ca ON ca.ID = c.category_id;