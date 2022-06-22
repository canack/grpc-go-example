package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/canack/grpc-example-go/services/grpc"
	"github.com/canack/grpc-example-go/services/order-service/database"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var servicePort = 50051

type server struct {
	pb.UnimplementedMicroserviceServer
}

func (s *server) Get(ctx context.Context, in *pb.RequestGet) (*pb.ReplyGet, error) {

	var tmpOrder database.Order

	tmpOrder.OrderUUID = in.GetUUID()

	response, err := tmpOrder.GetOrder()

	return &pb.ReplyGet{Data: response}, err
}

func (s *server) Create(ctx context.Context, in *pb.RequestCreate) (*pb.ReplyCreate, error) {

	var tmpOrder database.Order

	json.Unmarshal(in.GetData(), &tmpOrder)

	now := time.Now()
	tmpOrder.OrderUUID = uuid.NewString()
	tmpOrder.CreatedAt = now
	tmpOrder.UpdatedAt = now

	tmpOrder.Status = "New order"

	if tmpOrder.CustomerUUID == "" {
		return &pb.ReplyCreate{UUID: "", Status: false}, errors.New("you should put the customer's UUIDv4")
	}

	createStatus, err := tmpOrder.CreateOrder()

	if len(createStatus) > 1 {
		return &pb.ReplyCreate{UUID: tmpOrder.OrderUUID, Status: true}, nil
	}

	return &pb.ReplyCreate{UUID: "", Status: false}, err

}

func (s *server) ChangeStatus(ctx context.Context, in *pb.RequestChangeStatus) (*pb.ReplyChangeStatus, error) {
	UUID := in.GetUUID()
	NewStatus := in.GetNewStatus()

	var tmpOrder = database.Order{OrderUUID: UUID, Status: NewStatus}

	success, err := tmpOrder.ChangeStatus()

	if err != nil {
		return &pb.ReplyChangeStatus{Status: false}, err
	}

	if success {
		return &pb.ReplyChangeStatus{Status: true}, nil
	}
	return &pb.ReplyChangeStatus{Status: false}, nil

}

func (s *server) Update(ctx context.Context, in *pb.RequestUpdate) (*pb.ReplyUpdate, error) {
	var tmpOrder database.Order

	json.Unmarshal(in.GetData(), &tmpOrder)

	success, err := tmpOrder.UpdateOrder()
	if err != nil {
		return &pb.ReplyUpdate{Status: false}, err
	}
	if success {
		return &pb.ReplyUpdate{Status: true}, nil
	}

	return &pb.ReplyUpdate{Status: false}, nil

}

func (s *server) Delete(ctx context.Context, in *pb.RequestDelete) (*pb.ReplyDelete, error) {
	var tmpOrder = database.Order{OrderUUID: in.GetUUID()}

	success, err := tmpOrder.DeleteOrder()

	if err != nil {
		return &pb.ReplyDelete{Status: false}, err
	}
	if success {
		return &pb.ReplyDelete{Status: true}, nil
	}

	return &pb.ReplyDelete{Status: false}, nil

}

func (s *server) CreateTestData(ctx context.Context, in *pb.RequestTestData) (*pb.ReplyTestData, error) {
	err := database.CreateTestOrders()
	if err != nil {
		return &pb.ReplyTestData{}, err
	}
	return &pb.ReplyTestData{}, nil
}

func (s *server) DeleteTestData(ctx context.Context, in *pb.RequestTestData) (*pb.ReplyTestData, error) {
	err := database.DeleteTestOrders()
	if err != nil {
		return &pb.ReplyTestData{}, err
	}
	return &pb.ReplyTestData{}, nil
}

func StartGRPC() {

	if env := os.Getenv("SERVICE_PORT"); env != "" {
		portInt, err := strconv.Atoi(env)
		if err != nil {
			panic("Port isn't a integer")
		}
		servicePort = portInt
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", servicePort))
	if err != nil {
		log.Fatalf("TCP connection error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMicroserviceServer(s, &server{})
	log.Printf("Service wasn't registered: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("An error occurred when service starting: %v", err)
	}
}
