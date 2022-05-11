package model

// Office struck for db table - office

type Office struct {
	Code       int    `json:"code"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode int    `json:"postalcode"`
	Territory  string `json:"territory"`
}
