package database

import "time"

type Address struct {
	AddressLine string `json:"addressLine"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CityCode    int    `json:"cityCode"`
}

type Order struct {
	OrderUUID    string    `json:"orderUUID"`
	CustomerUUID string    `json:"customerUUID"`
	Status       string    `json:"status"`
	Quantity     int       `json:"quantity"`
	Price        float32   `json:"price"`
	Address      Address   `json:"address"`
	Product      Product   `json:"product"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Product struct {
	ProductUUID string `json:"productUUID"`
	ImageUrl    string `json:"imageUrl"`
	Name        string `json:"name"`
}
