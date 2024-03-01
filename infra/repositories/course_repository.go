package repositories

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/gstanleysilva/go-sqlc-example/infra/database/gen"
	"github.com/gstanleysilva/go-sqlc-example/internal/domain"
	"github.com/gstanleysilva/go-sqlc-example/pkg/sqlc"
)

type CourseRepository struct {
	TxHelper *sqlc.SQLCTxHelper
	Queries  *db.Queries
}

func NewCourseRepository(sqlConn *sql.DB) *CourseRepository {
	return &CourseRepository{
		TxHelper: sqlc.NewSQLCHelper(sqlConn),
		Queries:  db.New(sqlConn),
	}
}

func (r *CourseRepository) GetAll() ([]domain.Course, error) {
	result := make([]domain.Course, 0)

	courses, err := r.Queries.ListCourses(context.Background())
	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		result = append(result, domain.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description.String,
			Price:       course.Price,
			CategoryID:  course.CategoryID,
		})
	}

	return result, nil
}

func (r *CourseRepository) GetById(id string) (*domain.Course, error) {

	course, err := r.Queries.GetCourse(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &domain.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description.String,
		Price:       course.Price,
		CategoryID:  course.CategoryID,
	}, nil
}

func (r *CourseRepository) Create(course *domain.Course) (*domain.Course, error) {

	err := r.Queries.CreateCourse(context.Background(), db.CreateCourseParams{
		ID:          course.ID,
		Name:        course.Name,
		Description: sql.NullString{String: course.Description, Valid: true},
		Price:       course.Price,
		CategoryID:  course.CategoryID,
	})
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *CourseRepository) Update(id string, course *domain.Course) (*domain.Course, error) {

	err := r.Queries.UpdateCourse(context.Background(), db.UpdateCourseParams{
		ID:          id,
		Name:        course.Name,
		Description: sql.NullString{String: course.Description, Valid: true},
		Price:       course.Price,
		CategoryID:  course.CategoryID,
	})
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *CourseRepository) Delete(id string) error {
	return r.Queries.DeleteCourse(context.Background(), id)
}
