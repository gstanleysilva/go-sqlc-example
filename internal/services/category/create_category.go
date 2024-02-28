package category

import "github.com/gstanleysilva/go-sqlc-example/internal/domain"

type CreateCategoryInput struct {
	Name        string
	Description string
}

type CreateCategoryService struct {
	cr domain.CategoryRepository
}

func NewCreateCategoryService(cr domain.CategoryRepository) *CreateCategoryService {
	return &CreateCategoryService{cr: cr}
}

func (s *CreateCategoryService) Execute(input CreateCategoryInput) (*domain.Category, error) {
	category := domain.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	createdCourse, err := s.cr.Create(&category)
	if err != nil {
		return nil, err
	}

	return createdCourse, nil
}
