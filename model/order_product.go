package model

// Order_product struck for db table - order_product

type Order_product struct {
	Id          int     `json:"id"`
	OrderId     int     `json:"orderid"`
	Productcode int     `json:"productcode"`
	Qty         int     `json:"qty"`
	Priceeach   float32 `json:"priceeach"`
}
