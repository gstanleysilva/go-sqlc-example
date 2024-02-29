package main

import (
	"database/sql"
	"fmt"

	"github.com/gstanleysilva/go-sqlc-example/infra/repositories"
	"github.com/gstanleysilva/go-sqlc-example/internal/services/coursecategory"
)

func main() {
	sqlConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer sqlConn.Close()

	//CREATE CATEGORY
	// cr := repositories.NewCreateCategoryService(sqlConn)
	// getCategoryService := category.NewGetCategoryService(cr)

	// category, err := getCategoryService.Execute("1")
	// if err != nil {
	// 	panic(err)
	// }

	//CREATE COURSE WITH CATEGORY
	repo := repositories.NewCourseCategoryRepository(sqlConn)
	createCourseWithCategory := coursecategory.NewCreateCourseWithCategoryService(repo)
	course, category, err := createCourseWithCategory.Execute(coursecategory.CourseInput{
		Name:        "Course 1",
		Description: "Course 1 description",
		Price:       10.00,
	}, coursecategory.CategoryInput{
		Name:        "Category 1",
		Description: "Category 1 description",
	})
	if err != nil {
		panic(err)
	}

	// CREATE COURSE AND CATEGORY WITH UOW
	// uow := uow.NewUow(sqlConn)
	// uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
	// 	repo := repositories.NewCategoryRepository(sqlConn)
	// 	repo.Queries = db.New(tx)
	// 	return repo
	// })

	// uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
	// 	repo := repositories.NewCourseRepository(sqlConn)
	// 	repo.Queries = db.New(tx)
	// 	return repo
	// })

	// createCourseWithCategory := coursecategory.NewCreateCourseAndCategoryUowService(uow)
	// course, category, err := createCourseWithCategory.Execute(coursecategory.CourseInput{
	// 	Name:        "Course 2",
	// 	Description: "Course 2 description",
	// 	Price:       20.00,
	// }, coursecategory.CategoryInput{
	// 	Name:        "Category 2",
	// 	Description: "Category 2 description",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(course)
	fmt.Println(category)

	//GET COURSE WITH CATEGORY
	// repo := repositories.NewCourseCategoryRepository(sqlConn)
	// getCoursesWithCategory := coursecategory.NewGetCoursesWithCategory(repo)
	// report, err := getCoursesWithCategory.Execute()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, course := range report {
	// 	println(fmt.Sprintf("CourseID: %v, CourseName: %v, CategoryID: %v, CategoryName: %v", course.CourseID, course.CourseName, course.CategoryID, course.CategoryName))
	// }

}
