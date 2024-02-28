-- name: GetCoursesWithCategories :many
SELECT  c.id as CourseID,
        c.name as CourseName, 
        c.description as CourseDescription, 
        c.price as CoursePrice, 
        ct.id as CategoryID,
        ct.name as CategoryName,
        ct.description as CategoryDescription
FROM courses c
LEFT JOIN categories ct ON c.category_id = ct.id
GROUP BY c.id;