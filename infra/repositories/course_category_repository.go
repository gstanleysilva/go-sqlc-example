package repositories

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/gstanleysilva/go-sqlc-example/infra/database/gen"
	"github.com/gstanleysilva/go-sqlc-example/infra/sqlc"
	"github.com/gstanleysilva/go-sqlc-example/internal/domain"
)

type CourseCategoryRepository struct {
	TxHelper *sqlc.SQLCTxHelper
	Queries  *db.Queries
}

func NewCourseCategoryRepository(sqlConn *sql.DB) *CourseCategoryRepository {
	return &CourseCategoryRepository{
		TxHelper: sqlc.NewSQLCHelper(sqlConn),
		Queries:  db.New(sqlConn),
	}
}

func (r *CourseCategoryRepository) GetCoursesAndRepositories() ([]domain.CourseAndCategory, error) {
	result := make([]domain.CourseAndCategory, 0)

	values, err := r.Queries.GetCoursesWithCategories(context.Background())
	if err != nil {
		return result, err
	}

	for _, value := range values {
		result = append(result, domain.CourseAndCategory{
			CourseID:            value.Courseid,
			CourseName:          value.Coursename,
			CourseDescription:   value.Coursedescription.String,
			CoursePrice:         value.Courseprice,
			CategoryID:          value.Categoryid.String,
			CategoryName:        value.Categoryname.String,
			CategoryDescription: value.Categorydescription.String,
		})
	}

	return result, nil
}

func (r *CourseCategoryRepository) CreateCourseWithCategory(course *domain.Course, category *domain.Category) (*domain.CourseAndCategory, error) {
	ctx := context.Background()

	err := r.TxHelper.CallTx(context.Background(), func(q *db.Queries) error {
		_, err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          course.ID,
			Name:        course.Name,
			Description: sql.NullString{String: course.Description, Valid: true},
			Price:       course.Price,
			CategoryID:  category.ID,
		})
		return err
	})

	if err != nil {
		return nil, err
	}

	return &domain.CourseAndCategory{
		CourseID:            course.ID,
		CourseName:          course.Name,
		CourseDescription:   course.Description,
		CoursePrice:         course.Price,
		CategoryID:          category.ID,
		CategoryName:        category.Name,
		CategoryDescription: category.Description,
	}, nil
}
