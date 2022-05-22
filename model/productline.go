package model

// Productline struck for db table - productline

type Productline struct {
	Id         int    `json:"id"`
	Desclntext string `json:"desclntext"`
	Desclnhtml string `json:"desclnhtml"`
	Image      string `json:"image"`
}
