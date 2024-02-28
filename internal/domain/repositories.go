package domain

type CourseRepository interface {
	GetAll() ([]Course, error)
	GetById(id string) (*Course, error)
	Create(course *Course) (*Course, error)
	Update(id string, course *Course) (*Course, error)
	Delete(id string) error
}

type CategoryRepository interface {
	GetAll() ([]Category, error)
	GetById(id string) (*Category, error)
	Create(category *Category) (*Category, error)
	Update(id string, category *Category) (*Category, error)
	Delete(id string) error
}

type CourseCategoryRepository interface {
	CreateCourseWithCategory(course *Course, category *Category) (*CourseAndCategory, error)
	GetCoursesAndRepositories() ([]CourseAndCategory, error)
}
