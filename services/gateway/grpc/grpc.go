// This file contains gRPC communications

package grpc

import (
	"context"
	pb "github.com/canack/grpc-example-go/services/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

// Caution!!
// Addresses should change in docker-compose too.
// Otherwise, they will not be able to communicate with each other
var orderServiceAddress = "order:50051"
var customerServiceAddress = "customer:50052"

type services struct {
	// We have only 2 services
	// If more services would be available on future, we will define here
	CustomerService, OrderService pb.MicroserviceClient
}

var Service services

func connectCustomerService() {
	customerConn, err := grpc.Dial(customerServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	Service.CustomerService = pb.NewMicroserviceClient(customerConn)

}

func connectOrderService() {
	orderConn, err := grpc.Dial(orderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	Service.OrderService = pb.NewMicroserviceClient(orderConn)

}

func SetupServices() {

	if env := os.Getenv("ORDER_SERVICE_ADDRESS"); env != "" {
		orderServiceAddress = env
	}

	if env := os.Getenv("CUSTOMER_SERVICE_ADDRESS"); env != "" {
		customerServiceAddress = env
	}

	connectOrderService()
	connectCustomerService()
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
func SendGet(service pb.MicroserviceClient, payload string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := service.Get(ctx, &pb.RequestGet{UUID: payload})
	if err != nil {
		return []byte{}, err
	}

	return r.GetData(), nil
}

func SendCreate(service pb.MicroserviceClient, payload []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := service.Create(ctx, &pb.RequestCreate{Data: payload})

	if err != nil {
		return "", err
	}

	if r.GetStatus() {
		return r.GetUUID(), nil
	} else {
		return "", nil
	}
}

func SendDelete(service pb.MicroserviceClient, uuid string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := service.Delete(ctx, &pb.RequestDelete{UUID: uuid})
	if err != nil {
		return false, err
	}
	return r.GetStatus(), nil
}

func SendUpdate(service pb.MicroserviceClient, payload []byte) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := service.Update(ctx, &pb.RequestUpdate{Data: payload})
	if err != nil {
		return false, err
	}
	return r.GetStatus(), nil
}

func SendValidate(service pb.MicroserviceClient, uuid string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := service.Validate(ctx, &pb.RequestValidate{UUID: uuid})
	if err != nil {
		return false, err
	}
	return r.GetStatus(), nil
}

func SendChangeStatus(service pb.MicroserviceClient, uuid, newStatus string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := service.ChangeStatus(ctx, &pb.RequestChangeStatus{UUID: uuid, NewStatus: newStatus})
	if err != nil {
		return false, err
	}

	return r.GetStatus(), nil

}

func SendCreateTestData(service pb.MicroserviceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := service.CreateTestData(ctx, &pb.RequestTestData{})
	if err != nil {
		return err
	}
	return nil
}

func SendDeleteTestData(service pb.MicroserviceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := service.DeleteTestData(ctx, &pb.RequestTestData{})
	if err != nil {
		return err
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////
