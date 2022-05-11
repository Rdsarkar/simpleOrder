package model

// Order struck for db table - order

type Order struct {
	Id           int    `json:"id"`
	CustomerId   int    `json:"customerid"`
	OrderDate    string `json:"orderdate"`
	RequiredDate string `json:"requireddate"`
	ShippedDate  string `json:"shippeddate"`
	Status       int    `json:"status"`
	Comments     string `json:"comments"`
}
