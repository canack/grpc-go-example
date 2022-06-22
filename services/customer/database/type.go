package database

import "time"

type Address struct {
	AddressLine string `json:"addressLine"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CityCode    int    `json:"cityCode"`
}

type Customer struct {
	CustomerUUID string    `json:"customerUUID"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Address      Address   `json:"address"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
