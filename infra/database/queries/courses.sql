-- name: CreateCourse :exec
insert into courses (id, name, description, category_id, price)
VALUES (?, ?, ?, ?, ?);

-- name: GetCourse :one
SELECT *
FROM courses
WHERE id = ?;

-- name: ListCourses :many
SELECT *
FROM courses
ORDER BY id;

-- name: UpdateCourse :exec
UPDATE courses
SET name = ?,
    description = ?,
    category_id = ?,
    price = ?
WHERE id = ?;

-- name: DeleteCourse :exec
DELETE
FROM courses
WHERE id = ?;