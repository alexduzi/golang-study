-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (ID, name, description) 
VALUES (?,?,?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ?
WHERE ID = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE ID = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id, price)
VALUES (?, ?, ?, ?, ?);

-- name: ListCourses :many
SELECT c.*, ca.name as category_name
FROM courses c
JOIN categories ca ON ca.ID = c.category_id;