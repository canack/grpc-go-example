package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/canack/grpc-example-go/services/customer-service/database"
	pb "github.com/canack/grpc-example-go/services/grpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var servicePort = 50052

type server struct {
	pb.UnimplementedMicroserviceServer
}

func (s *server) Get(ctx context.Context, in *pb.RequestGet) (*pb.ReplyGet, error) {

	var tmpCustomer database.Customer

	tmpCustomer.CustomerUUID = in.GetUUID()

	response, err := tmpCustomer.GetCustomer()

	return &pb.ReplyGet{Data: response}, err
}

func (s *server) Create(ctx context.Context, in *pb.RequestCreate) (*pb.ReplyCreate, error) {

	var tmpCustomer database.Customer

	json.Unmarshal(in.GetData(), &tmpCustomer)

	now := time.Now()
	tmpCustomer.CustomerUUID = uuid.NewString()
	tmpCustomer.CreatedAt = now
	tmpCustomer.UpdatedAt = now

	createStatus, err := tmpCustomer.CreateCustomer()

	if len(createStatus) > 1 {
		return &pb.ReplyCreate{UUID: tmpCustomer.CustomerUUID, Status: true}, nil
	}

	return &pb.ReplyCreate{UUID: "", Status: false}, err

}

func (s *server) Update(ctx context.Context, in *pb.RequestUpdate) (*pb.ReplyUpdate, error) {
	var tmpCustomer database.Customer

	json.Unmarshal(in.GetData(), &tmpCustomer)

	success, err := tmpCustomer.UpdateCustomer()
	if err != nil {
		return &pb.ReplyUpdate{Status: false}, err
	}
	if success {
		return &pb.ReplyUpdate{Status: true}, nil

	}

	return &pb.ReplyUpdate{Status: false}, nil

}

func (s *server) Delete(ctx context.Context, in *pb.RequestDelete) (*pb.ReplyDelete, error) {

	var tmpCustomer = database.Customer{CustomerUUID: in.GetUUID()}

	success, err := tmpCustomer.DeleteCustomer()

	if err != nil {
		return &pb.ReplyDelete{Status: false}, err
	}
	if success {
		return &pb.ReplyDelete{Status: true}, nil
	}

	return &pb.ReplyDelete{Status: false}, nil

}

func (s *server) Validate(ctx context.Context, in *pb.RequestValidate) (*pb.ReplyValidate, error) {
	var tmpCustomer database.Customer
	tmpCustomer.CustomerUUID = in.GetUUID()
	status, err := tmpCustomer.CheckCustomer()
	if err != nil {
		return &pb.ReplyValidate{Status: false}, err
	}
	return &pb.ReplyValidate{Status: status}, nil
}

func (s *server) CreateTestData(ctx context.Context, in *pb.RequestTestData) (*pb.ReplyTestData, error) {
	err := database.CreateTestCustomers()
	if err != nil {
		return &pb.ReplyTestData{}, err
	}
	return &pb.ReplyTestData{}, nil
}

func (s *server) DeleteTestData(ctx context.Context, in *pb.RequestTestData) (*pb.ReplyTestData, error) {
	err := database.DeleteTestCustomers()
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
