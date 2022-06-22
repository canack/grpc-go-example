package types

type CustomerRequestCreate struct {
	Name    string  `json:"name" example:"John Doe"`
	Email   string  `json:"email" example:"johndoe@example.com"`
	Address Address `json:"address" `
}

type CustomerRequestUpdate struct {
	Name    string  `json:"name" example:"John Doe"`
	Email   string  `json:"email" example:"johndoe@example.com"`
	Address Address `json:"address"`
}

type CustomerGet struct {
	UUIDv4 string `example:"190edd4b-a89c-4f74-b7e0-256645fd0373"`
}

type OrderRequestCreate struct {
	CustomerUUID string  `json:"customerUUID" example:"190edd4b-a89c-4f74-b7e0-256645fd0373"`
	Status       string  `json:"status" example:"New order"`
	Quantity     int     `json:"quantity" example:"2"`
	Price        float32 `json:"price" example:"49.99"`
	Address      Address `json:"address"`
	Product      Product `json:"product"`
}

type OrderRequestUpdate struct {
	Status   string  `json:"status" example:"Preparing"`
	Quantity int     `json:"quantity" example:"3"`
	Price    float32 `json:"price" example:"69.99"`
	Address  Address `json:"address"`
	Product  Product `json:"product"`
}

type OrderRequestUpdateStatus struct {
	Status string `json:"status" example:"Shipped"`
}
