package vacancy

import "time"

type VacancyCreateForm struct {
	Role     string
	Company  string
	Sphere   string
	Salary   string
	Location string
	Email    string
}

type Vacancy struct {
	Id        int
	Role      string
	Company   string
	Sphere    string
	Salary    string
	Location  string
	Email     string
	CreatedAt time.Time
}
