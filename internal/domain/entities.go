package domain

type Course struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryID  string
}

type Category struct {
	ID          string
	Name        string
	Description string
}

type CourseAndCategory struct {
	CourseID            string
	CourseName          string
	CourseDescription   string
	CoursePrice         float64
	CategoryID          string
	CategoryName        string
	CategoryDescription string
}
