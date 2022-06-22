package handler

import (
	"github.com/canack/grpc-example-go/services/gateway-service/grpc"
	"github.com/labstack/echo/v4"
	"net/http"
)

var testCreated bool

// @Summary      Creates test database records for orders and customers
// @Description  Creates test database records for orders and customers
// @Tags         Database operations
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router       /generate [get]
func CreateTestData(c echo.Context) error {
	if testCreated {
		return c.String(http.StatusOK, "Test records is already created!")
	}

	err := grpc.SendCreateTestData(grpc.Service.CustomerService)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = grpc.SendCreateTestData(grpc.Service.OrderService)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	testCreated = true
	return c.String(http.StatusOK, "Test records successfully created")
}

// @Summary      Deletes test database records for orders and customers
// @Description  Deletes test database records for orders and customers
// @Tags         Database operations
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router       /reset [get]
func DeleteTestData(c echo.Context) error {
	err := grpc.SendDeleteTestData(grpc.Service.CustomerService)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = grpc.SendDeleteTestData(grpc.Service.OrderService)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	testCreated = false
	return c.String(http.StatusOK, "All records have been deleted")
}
