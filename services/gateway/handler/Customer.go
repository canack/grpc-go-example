// Processing endpoint requests here

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/canack/grpc-example-go/services/gateway-service/grpc"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary      Returns all customers
// @Description  Returns all customers as json array
// @Tags         customer
// @Accept       json
// @Produce      json
// @Success      200  {object}  []types.Customer
// @Failure      400  {object}  string
// @Router       /customer/ [get]
func CustomerGet(c echo.Context) error {
	customerUUID := c.Param("customer")

	responseBytes, err := grpc.SendGet(grpc.Service.CustomerService, customerUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("An problem occured: %v", err))
	}

	var responseCustomer []Customer

	json.Unmarshal(responseBytes, &responseCustomer)

	return c.JSON(http.StatusOK, responseCustomer)
}

// @Summary      Returns customer data with given UUIDv4
// @Description  Returns customer data as json with given UUIDv4
// @Tags         customer
// @Param        UUIDv4 path string true "Customer UUIDv4"
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Customer
// @Failure      400  {object}  string
// @Router       /customer/{UUIDv4} [get]
func CustomerGetUUID(c echo.Context) error {
	customerUUID := c.Param("customer")

	responseBytes, err := grpc.SendGet(grpc.Service.CustomerService, customerUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("An problem occured: %v", err))
	}

	var responseCustomer []Customer

	json.Unmarshal(responseBytes, &responseCustomer)

	return c.JSON(http.StatusOK, responseCustomer)
}

// @Summary      Creates a new customer
// @Description  Creates a new customer and returns the customer's UUIDv4
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param customer body types.CustomerRequestCreate true "Customer info"
// @Success      201  {object}  string
// @Failure      400  {object}  string
// @Failure      502  {object}  string
// @Router       /customer/ [post]
func CustomerCreate(c echo.Context) error {
	customerRequest := new(Customer)

	if err := c.Bind(customerRequest); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Request can't parsing."))
	}

	customerBytes, err := json.Marshal(customerRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Request can't parsing: %v", err.Error()))
	}

	CustomerUUID, err := grpc.SendCreate(grpc.Service.CustomerService, customerBytes)
	if err != nil {
		return c.String(http.StatusBadGateway, fmt.Sprintf("%v", err))
	}

	return c.String(http.StatusCreated, CustomerUUID)

}

// @Summary      Deletes the customer
// @Description  Deletes the customer and returns boolean
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Customer UUIDv4"
// @Success      200  {object} bool
// @Failure		 400 {object} string
// @Router       /customer/{UUIDv4} [delete]
func CustomerDelete(c echo.Context) error {
	customerUUID := c.Param("customer")
	var tempCustomer Customer

	tempCustomer.CustomerUUID = customerUUID

	status, err := grpc.SendDelete(grpc.Service.CustomerService, customerUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))

}

// @Summary      Updates customer's info
// @Description  Updates customer's info and returns boolean
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Customer UUIDv4"
// @Param customer body types.CustomerRequestUpdate true "Customer info"
// @Success      200  {object}  bool
// @Failure      400 {object} string
// @Router       /customer/{UUIDv4} [put]
func CustomerUpdate(c echo.Context) error {
	customerUUID := c.Param("customer")
	changeRequest := new(Customer)

	if err := c.Bind(changeRequest); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Request can't parsing."))
	}

	changeRequest.CustomerUUID = customerUUID

	customerBytes, err := json.Marshal(changeRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request can't parsing.")
	}

	status, err := grpc.SendUpdate(grpc.Service.CustomerService, customerBytes)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request can't parsing.")
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))

}

// @Summary      Checks if the customer is in the database
// @Description  Checks if the customer is in the database and returns boolean
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Customer UUIDv4"
// @Success      200  {object}  bool
// @Failure      400 {object} string
// @Router       /customer/validate/{UUIDv4} [get]
func CustomerValidate(c echo.Context) error {
	customerUUID := c.Param("customer")

	status, err := grpc.SendValidate(grpc.Service.CustomerService, customerUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request can't parsing.")
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))
	
}
