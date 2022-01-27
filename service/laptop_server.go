package service

import (
	"context"
	"log"

	"github.com/giavudangle/go-grpc/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopRequest, error) {
	laptop := req.Laptop
	log.Printf("receive a create-laptop request with id: %w", laptop.Id)
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil,
				status.Error(codes.InvalidArgument, "Laptop ID is not a valid UUID %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Error(codes.Internal, "cannot generate a new laptop ID: %w", err)
		}

		laptop.Id = id.String()
	}

}
