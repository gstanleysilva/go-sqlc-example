package coursecategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/gstanleysilva/go-sqlc-example/internal/domain"
	"github.com/gstanleysilva/go-sqlc-example/pkg/uow"
)

type CreateCourseAndCategoryUowService struct {
	Uow uow.UowInterface
}

func NewCreateCourseAndCategoryUowService(uow uow.UowInterface) *CreateCourseAndCategoryUowService {
	return &CreateCourseAndCategoryUowService{
		Uow: uow,
	}
}

func (s *CreateCourseAndCategoryUowService) Execute(courseArgs CourseInput, categoryArgs CategoryInput) (*domain.Course, *domain.Category, error) {
	categoryId := uuid.NewString()
	ctx := context.Background()

	var createdCategory *domain.Category
	var createdCourse *domain.Course

	err := s.Uow.Do(ctx, func(uow uow.UowInterface) error {
		var err error

		categoryRepo := s.getCategoryRepository(ctx)
		courseRepo := s.getCourseRepository(ctx)

		createdCategory, err = categoryRepo.Create(&domain.Category{
			ID:          categoryId,
			Name:        categoryArgs.Name,
			Description: categoryArgs.Description,
		})
		if err != nil {
			return err
		}

		createdCourse, err = courseRepo.Create(&domain.Course{
			ID:          uuid.NewString(),
			Name:        courseArgs.Name,
			Description: courseArgs.Description,
			CategoryID:  categoryId,
			Price:       courseArgs.Price,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return createdCourse, createdCategory, nil
}

func (s *CreateCourseAndCategoryUowService) getCourseRepository(ctx context.Context) domain.CourseRepository {
	repo, err := s.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(domain.CourseRepository)
}

func (s *CreateCourseAndCategoryUowService) getCategoryRepository(ctx context.Context) domain.CategoryRepository {
	repo, err := s.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(domain.CategoryRepository)
}
