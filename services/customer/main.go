// This file contains main function of service
package main

import (
	db "github.com/canack/grpc-example-go/services/customer-service/database"
	"github.com/canack/grpc-example-go/services/customer-service/grpc"
)

func main() {
	db.DBStart()
	db.CreateTables()
	grpc.StartGRPC()
}
