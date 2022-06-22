// Processing endpoint requests here

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/canack/grpc-example-go/services/gateway-service/grpc"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary      Returns all orders
// @Description  Returns all orders as json array
// @Tags         order
// @Accept       json
// @Produce      json
// @Success      200  {object}  []types.Order
// @Failure      400  {object}  string
// @Router       /order/ [get]
func OrderGet(c echo.Context) error {
	orderUUID := c.Param("order")

	responseBytes, err := grpc.SendGet(grpc.Service.OrderService, orderUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("An problem occured: %v", err))
	}

	var responseOrder []Order

	json.Unmarshal(responseBytes, &responseOrder)

	return c.JSON(http.StatusOK, responseOrder)
}

// @Summary      Returns order data with given UUIDv4
// @Description  Returns order data as json with given UUIDv4
// @Tags         order
// @Param        UUIDv4 path string true "Order UUIDv4"
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Order
// @Failure      400  {object}  string
// @Router       /order/{UUIDv4} [get]
func OrderGetUUID(c echo.Context) error {
	orderUUID := c.Param("order")

	responseBytes, err := grpc.SendGet(grpc.Service.OrderService, orderUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("An problem occured: %v", err))
	}

	var responseOrder []Order

	json.Unmarshal(responseBytes, &responseOrder)

	return c.JSON(http.StatusOK, responseOrder)
}

// @Summary      Creates a new order
// @Description  Creates a new order and returns the order's UUIDv4
// @Tags         order
// @Accept       json
// @Produce      json
// @Param order body types.OrderRequestCreate true "Order info"
// @Success      201  {object}  string
// @Failure      400  {object}  string
// @Failure      502  {object}  string
// @Router       /order/ [post]
func OrderCreate(c echo.Context) error {
	orderRequest := new(Order)

	if err := c.Bind(orderRequest); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Request can't parsing."))
	}

	orderBytes, err := json.Marshal(orderRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Request can't parsing: %v", err.Error()))
	}

	OrderUUID, err := grpc.SendCreate(grpc.Service.OrderService, orderBytes)
	if err != nil {
		return c.String(http.StatusBadGateway, fmt.Sprintf("%v", err))
	}

	return c.String(http.StatusCreated, OrderUUID)

}

// @Summary      Deletes the order
// @Description  Deletes the order and returns boolean
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Order UUIDv4"
// @Success      200  {object} bool
// @Failure		 400 {object} string
// @Router       /order/{UUIDv4} [delete]
func OrderDelete(c echo.Context) error {
	orderUUID := c.Param("order")
	var tempOrder Order

	tempOrder.OrderUUID = orderUUID

	status, err := grpc.SendDelete(grpc.Service.OrderService, orderUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))

}

// @Summary      Updates order's status
// @Description  Updates order's status and returns boolean
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Order UUIDv4"
// @Param order body types.OrderRequestUpdateStatus true "Order status"
// @Success      200  {object}  bool
// @Failure      400 {object} string
// @Router       /order/{UUIDv4} [patch]
func OrderChangeStatus(c echo.Context) error {
	orderUUID := c.Param("order")

	changeRequest := new(Order)

	if err := c.Bind(changeRequest); err != nil {
		return c.String(http.StatusBadRequest, "Request can't parsing.")
	}

	status, err := grpc.SendChangeStatus(grpc.Service.OrderService, orderUUID, changeRequest.Status)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))

}

// @Summary      Updates order's info
// @Description  Updates order's info and returns boolean
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        UUIDv4 path string true "Order UUIDv4"
// @Param order body types.OrderRequestUpdate true "Order info"
// @Success      200  {object}  bool
// @Failure      400 {object} string
// @Router       /order/{UUIDv4} [put]
func OrderUpdate(c echo.Context) error {
	orderUUID := c.Param("order")
	changeRequest := new(Order)

	if err := c.Bind(changeRequest); err != nil {
		return c.String(http.StatusBadRequest, "Request can't parsing.")
	}

	changeRequest.OrderUUID = orderUUID

	orderBytes, err := json.Marshal(changeRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request can't parsing.")
	}

	status, err := grpc.SendUpdate(grpc.Service.OrderService, orderBytes)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request can't parsing.")
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", status))

}
