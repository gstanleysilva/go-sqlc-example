package coursecategory

import (
	"github.com/google/uuid"
	"github.com/gstanleysilva/go-sqlc-example/internal/domain"
)

type CourseInput struct {
	Name        string
	Description string
	Price       float64
}

type CategoryInput struct {
	Name        string
	Description string
}

type CreateCourseWithCategoryService struct {
	repo domain.CourseCategoryRepository
}

func NewCreateCourseWithCategoryService(repo domain.CourseCategoryRepository) *CreateCourseWithCategoryService {
	return &CreateCourseWithCategoryService{
		repo: repo,
	}
}

func (s *CreateCourseWithCategoryService) Execute(courseArgs CourseInput, categoryArgs CategoryInput) (*domain.Course, *domain.Category, error) {
	categoryId := uuid.NewString()

	result, err := s.repo.CreateCourseWithCategory(&domain.Course{
		ID:          uuid.NewString(),
		Name:        courseArgs.Name,
		Description: courseArgs.Description,
		Price:       courseArgs.Price,
		CategoryID:  categoryId,
	}, &domain.Category{
		ID:          categoryId,
		Name:        categoryArgs.Name,
		Description: categoryArgs.Description,
	})
	if err != nil {
		return nil, nil, err
	}

	course := domain.Course{
		ID:          result.CourseID,
		Name:        result.CourseName,
		Description: result.CourseDescription,
		Price:       result.CoursePrice,
		CategoryID:  result.CategoryID,
	}

	category := domain.Category{
		ID:          result.CategoryID,
		Name:        result.CategoryName,
		Description: result.CategoryDescription,
	}

	return &course, &category, nil
}
