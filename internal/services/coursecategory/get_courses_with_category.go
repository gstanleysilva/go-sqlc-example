package coursecategory

import "github.com/gstanleysilva/go-sqlc-example/internal/domain"

type GetCoursesWithCategory struct {
	repo domain.CourseCategoryRepository
}

func NewGetCoursesWithCategory(repo domain.CourseCategoryRepository) *GetCoursesWithCategory {
	return &GetCoursesWithCategory{
		repo: repo,
	}
}

func (s *GetCoursesWithCategory) Execute() ([]domain.CourseAndCategory, error) {

	entries, err := s.repo.GetCoursesAndRepositories()
	if err != nil {
		return []domain.CourseAndCategory{}, err
	}

	return entries, nil
}
