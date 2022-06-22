package types

import "time"

// General

type Address struct {
	AddressLine string `json:"addressLine" example:"x street, y apartment, number:5/7"`
	City        string `json:"city" example:"New york"`
	Country     string `json:"country" example:"United States"`
	CityCode    int    `json:"cityCode" example:"45"`
}

////////////////

// Order

type Order struct {
	OrderUUID    string    `json:"orderUUID"`
	CustomerUUID string    `json:"customerUUID"`
	Status       string    `json:"status"`
	Quantity     int       `json:"quantity"`
	Price        float32   `json:"price"`
	Address      Address   `json:"address"`
	Product      Product   `json:"product"`
	CreatedAt    time.Time `json:"createdAt" example:"2022-05-23T09:45:01.675884703Z"`
	UpdatedAt    time.Time `json:"updatedAt" example:"2022-05-23T09:55:04.995285414Z"`
}

type Product struct {
	ProductUUID string `json:"productUUID" example:"455avd4b-x22z-2554-z88m-0amz8552"`
	ImageUrl    string `json:"imageUrl" example:"https://example.com/product.jpg"`
	Name        string `json:"name" example:"Toy car"`
}

////////////////

// Customer

type Customer struct {
	CustomerUUID string    `json:"customerUUID"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Address      Address   `json:"address"`
	CreatedAt    time.Time `json:"createdAt" example:"2022-05-23T09:45:01.675884703Z"`
	UpdatedAt    time.Time `json:"updatedAt" example:"2022-05-23T09:55:04.995285414Z"`
}
