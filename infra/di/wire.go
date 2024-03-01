//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/gstanleysilva/go-sqlc-example/infra/repositories"
	"github.com/gstanleysilva/go-sqlc-example/internal/domain"
	"github.com/gstanleysilva/go-sqlc-example/internal/services/category"
	"github.com/gstanleysilva/go-sqlc-example/internal/services/coursecategory"
)

var setRepositoryDependency = wire.NewSet(
	repositories.NewCourseCategoryRepository,
	wire.Bind(new(domain.CourseCategoryRepository), new(*repositories.CourseCategoryRepository)),
	repositories.NewCategoryRepository,
	wire.Bind(new(domain.CategoryRepository), new(*repositories.CategoryRepository)),
)

func NewGetCategoryService(db *sql.DB) *category.GetCategoryService {
	wire.Build(
		setRepositoryDependency,
		category.NewGetCategoryService,
	)
	return &category.GetCategoryService{}
}

func NewCreateCategoryService(db *sql.DB) *category.CreateCategoryService {
	wire.Build(
		setRepositoryDependency,
		category.NewCreateCategoryService,
	)
	return &category.CreateCategoryService{}
}

func NewCreateCourseAndCategoryUowService(db *sql.DB) *coursecategory.CreateCourseAndCategoryUowService {
	wire.Build(
		coursecategory.NewCourseAndCategoryUow,
		coursecategory.NewCreateCourseAndCategoryUowService,
	)
	return &coursecategory.CreateCourseAndCategoryUowService{}
}

func NewGetCoursesWithCategoryService(db *sql.DB) *coursecategory.GetCoursesWithCategoryService {
	wire.Build(
		setRepositoryDependency,
		coursecategory.NewGetCoursesWithCategoryService,
	)
	return &coursecategory.GetCoursesWithCategoryService{}
}

func NewCreateCourseAndCategoryService(db *sql.DB) *coursecategory.CreateCourseWithCategoryService {
	wire.Build(
		setRepositoryDependency,
		coursecategory.NewCreateCourseWithCategoryService,
	)
	return &coursecategory.CreateCourseWithCategoryService{}
}
