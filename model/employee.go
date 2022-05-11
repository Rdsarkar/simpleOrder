package model

// Employee struck for db table - employee
type Employee struct {
	Id         int    `json:"id"`
	Officecode int    `json:"officecode"`
	Reportsto  int    `json:"reportsto"`
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Extension  string `json:"extension"`
	Email      string `json:"email"`
	Jobtitle   string `json:"jobtitle"`
}
