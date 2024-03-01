package category

import (
	"github.com/gstanleysilva/go-sqlc-example/internal/domain"
)

type GetCategoryService struct {
	cr domain.CategoryRepository
}

func NewGetCategoryService(cr domain.CategoryRepository) *GetCategoryService {
	return &GetCategoryService{cr: cr}
}

func (s GetCategoryService) Execute(id string) (domain.Category, error) {

	category, err := s.cr.GetById(id)
	if err != nil {
		return domain.Category{}, err
	}

	return *category, nil
}
