package main

import (
	"database/sql"
	"fmt"

	"github.com/gstanleysilva/go-sqlc-example/infra/di"
	"github.com/gstanleysilva/go-sqlc-example/internal/services/coursecategory"
)

func main() {
	sqlConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer sqlConn.Close()

	//CREATE CATEGORY
	getCategoryService := di.NewGetCategoryService(sqlConn)
	category, err := getCategoryService.Execute("1")
	if err == nil {
		fmt.Println(category)
	}

	//CREATE COURSE WITH CATEGORY
	createCourseWithCategory := di.NewCreateCourseAndCategoryService(sqlConn)
	_, _, err = createCourseWithCategory.Execute(coursecategory.CourseInput{
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
	createCourseWithCategoryUow := di.NewCreateCourseAndCategoryUowService(sqlConn)
	_, _, err = createCourseWithCategoryUow.Execute(coursecategory.CourseInput{
		Name:        "Course 2",
		Description: "Course 2 description",
		Price:       20.00,
	}, coursecategory.CategoryInput{
		Name:        "Category 2",
		Description: "Category 2 description",
	})
	if err != nil {
		panic(err)
	}

	//GET COURSE WITH CATEGORY
	getCoursesWithCategory := di.NewGetCoursesWithCategoryService(sqlConn)
	report, err := getCoursesWithCategory.Execute()
	if err != nil {
		panic(err)
	}

	for _, course := range report {
		println(fmt.Sprintf("CourseID: %v, CourseName: %v, CategoryID: %v, CategoryName: %v", course.CourseID, course.CourseName, course.CategoryID, course.CategoryName))
	}

}
