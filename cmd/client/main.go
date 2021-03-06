package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/giavudangle/go-grpc/pb"
	"github.com/giavudangle/go-grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func createLaptop(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	laptop.Id = "0e434306-d2aa-4c4e-98b5-82e2a6643f34"
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	log.Print("created laptop with id: %w", res.Id)
}

func SearchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Print("search filter: ", filter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchLaptopRequest{
		Filter: filter,
	}
	stream, err := laptopClient.SearchLaptop(ctx, req)

	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		laptop := res.GetLaptop()

		// Print laptop information
		log.Print("===Laptop Information===\n")
		log.Print("ID: ", laptop.GetId()+"\n")
		log.Print("Brand: ", laptop.GetBrand()+"\n")
		log.Print("Name: ", laptop.GetName()+"\n")
		log.Print("CPU Cores: ", laptop.GetCpu().GetNumberCores(), "\n")
		log.Print("CPU MinGhz: ", laptop.GetCpu().GetMinGhz(), "\n")
		log.Print("RAM: ", laptop.GetRam().GetValue(), laptop.GetRam().GetUnit(), "\n")
		log.Print("Price: ", laptop.GetPriceUsd(), " USD")

	}
}

func main() {
	fmt.Println("client")
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	for i := 0; i < 10; i++ {
		createLaptop(laptopClient)
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Unit:  pb.Memory_GIGABYTE,
			Value: 8,
		},
	}

	SearchLaptop(laptopClient, filter)

}
