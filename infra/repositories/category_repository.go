package repositories

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/gstanleysilva/go-sqlc-example/infra/database/gen"
	"github.com/gstanleysilva/go-sqlc-example/infra/sqlc"
	"github.com/gstanleysilva/go-sqlc-example/internal/domain"
)

type CategoryRepository struct {
	TxHelper *sqlc.SQLCTxHelper
	Queries  *db.Queries
}

func NewCreateCategoryService(sqlConn *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		TxHelper: sqlc.NewSQLCHelper(sqlConn),
		Queries:  db.New(sqlConn),
	}
}

func (r *CategoryRepository) GetAll() ([]domain.Category, error) {
	result := make([]domain.Category, 0)

	categories, err := r.Queries.ListCategories(context.Background())
	if err != nil {
		return []domain.Category{}, err
	}

	for _, cat := range categories {
		result = append(result, domain.Category(cat))
	}

	return result, nil
}

func (r *CategoryRepository) GetById(id string) (*domain.Category, error) {

	category, err := r.Queries.GetCategory(context.Background(), id)
	if err != nil {
		return nil, err
	}

	result := domain.Category(category)

	return &result, nil
}

func (r *CategoryRepository) Create(category *domain.Category) (*domain.Category, error) {

	_, err := r.Queries.CreateCategory(context.Background(), db.CreateCategoryParams(*category))
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(id string, category *domain.Category) (*domain.Category, error) {

	err := r.Queries.UpdateCategory(context.Background(), db.UpdateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	})
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Delete(id string) error {
	return r.Queries.DeleteCategory(context.Background(), id)
}
