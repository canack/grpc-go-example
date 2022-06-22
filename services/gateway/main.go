// This file contains main function.
// This file responsible for setting up endpoints and register gRPC services.
package main

import (
	_ "github.com/canack/grpc-example-go/services/gateway-service/api"
	"github.com/canack/grpc-example-go/services/gateway-service/grpc"
	ep "github.com/canack/grpc-example-go/services/gateway-service/handler"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	"net/http"
)

var serverAddress = ":3434"
var e *echo.Echo

// @title Customer & Order gRPC services
// @version 1.0
// @description Interactive documentation for Order & Customer services
func startServer() {
	e = echo.New()

	registerEndpoints()

	e.Logger.Fatal(e.Start(serverAddress))
}

func registerEndpoints() {
	// For order
	e.GET("/order/:order", ep.OrderGetUUID)
	e.GET("/order*", ep.OrderGet)
	e.POST("/order*", ep.OrderCreate)
	e.DELETE("/order/:order", ep.OrderDelete)
	e.PATCH("/order/:order", ep.OrderChangeStatus)
	e.PUT("/order/:order", ep.OrderUpdate)
	// For order

	// For customer
	e.GET("/customer/:customer", ep.CustomerGetUUID)
	e.GET("/customer/validate/:customer", ep.CustomerValidate)
	e.GET("/customer*", ep.CustomerGet)
	e.POST("/customer*", ep.CustomerCreate)
	e.DELETE("/customer/:customer", ep.CustomerDelete)
	e.PUT("/customer/:customer", ep.CustomerUpdate)
	// For customer

	// For test
	e.GET("/generate", ep.CreateTestData)
	e.GET("/reset", ep.DeleteTestData)
	// For test

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1>You're on main page</h1>
		<br><h2>For usage/examples:
		<a href='https://github.com/canack/grpc-example-go'>https://github.com/canack/grpc-example-go</a></h2>`)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

}

func main() {
	grpc.SetupServices()
	startServer()
}
