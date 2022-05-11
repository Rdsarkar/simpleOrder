package model

// Customer struck for db table - customer
type Customer struct {
	Id                  int     `json:"id"`
	SalesrepEmployeeNum int     `json:"salesrepemployeenum"`
	Name                string  `json:"name"`
	Lastname            string  `json:"lastname"`
	Firstname           string  `json:"firstname"`
	Phone               string  `json:"phone"`
	Address1            string  `json:"address1"`
	Address2            string  `json:"address2"`
	City                string  `json:"city"`
	State               string  `json:"state"`
	PostalCode          string  `json:"postalcode"`
	Country             string  `json:"country"`
	Creditlimit         float32 `json:"creditlimit"`
}
