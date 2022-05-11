package model

// Payment struck for db table - payment

type Payment struct {
	Checknum    string  `json:"checknum"`
	Customerid  int     `json:"customerid"`
	Paymentdate string  `json:"paymentdate"`
	Amount      float32 `json:"amount"`
}
