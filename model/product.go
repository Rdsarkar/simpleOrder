package model

// Product struck for db table - product

type Product struct {
	Code           int     `json:"code"`
	Productlineid  int     `json:"productlineid"`
	Name           string  `json:"name"`
	Scale          int     `json:"scale"`
	Vendor         string  `json:"vendor"`
	Pdtdescription string  `json:"pdtdescription"`
	Qtyinstorck    int     `json:"qtyinstock"`
	Buyprice       float32 `json:"buyprice"`
	Msrp           string  `json:"msrp"`
}
