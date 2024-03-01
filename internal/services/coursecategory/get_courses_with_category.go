package coursecategory

import "github.com/gstanleysilva/go-sqlc-example/internal/domain"

type GetCoursesWithCategoryService struct {
	repo domain.CourseCategoryRepository
}

func NewGetCoursesWithCategoryService(repo domain.CourseCategoryRepository) *GetCoursesWithCategoryService {
	return &GetCoursesWithCategoryService{
		repo: repo,
	}
}

func (s *GetCoursesWithCategoryService) Execute() ([]domain.CourseAndCategory, error) {

	entries, err := s.repo.GetCoursesAndRepositories()
	if err != nil {
		return []domain.CourseAndCategory{}, err
	}

	return entries, nil
}
