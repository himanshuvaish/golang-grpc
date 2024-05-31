package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/himanshuvaish/golang-grpc/pb"
	"google.golang.org/grpc"
)

// server is used to implement driver.DriverServiceServer.
type server struct {
	driverStore map[int32]*pb.Driver
}

// NewServer creates a new DriverServiceServer.
func NewServer() *server {
	drivers := map[int32]*pb.Driver{
		1: &pb.Driver{Id: 1, Name: "Lewis Hamilton", Wins: 98, Poles: 100},
		2: &pb.Driver{Id: 2, Name: "Sebastian Vettel", Wins: 53, Poles: 57},
		3: &pb.Driver{Id: 3, Name: "Max Verstappen", Wins: 22, Poles: 13},
		4: &pb.Driver{Id: 4, Name: "Valtteri Bottas", Wins: 10, Poles: 17},
	}
	return &server{
		driverStore: drivers,
	}
}

func (s *server) AddDriver(ctx context.Context, req *pb.AddDriverRequest) (*pb.AddDriverResponse, error) {
	driver := req.GetDriver()
	s.driverStore[driver.GetId()] = driver
	return &pb.AddDriverResponse{Driver: driver}, nil
}

func (s *server) GetDriver(ctx context.Context, req *pb.GetDriverRequest) (*pb.GetDriverResponse, error) {
	id := req.GetId()
	driver, ok := s.driverStore[id]
	if !ok {
		return nil, fmt.Errorf("driver with ID %d not found", id)
	}
	return &pb.GetDriverResponse{Driver: driver}, nil
}

func (s *server) UpdateDriver(ctx context.Context, req *pb.UpdateDriverRequest) (*pb.UpdateDriverResponse, error) {
	id := req.GetId()
	driver, ok := s.driverStore[id]
	if !ok {
		return nil, fmt.Errorf("driver with ID %d not found", id)
	}

	if wins := req.GetWins(); wins > 0 {
		driver.Wins = wins
	}

	if poles := req.GetPoles(); poles > 0 {
		driver.Poles = poles
	}

	return &pb.UpdateDriverResponse{Driver: driver}, nil
}

func (s *server) DeleteDriver(ctx context.Context, req *pb.DeleteDriverRequest) (*pb.DeleteDriverResponse, error) {
	id := req.GetId()
	_, ok := s.driverStore[id]
	if !ok {
		return nil, fmt.Errorf("driver with ID %d not found", id)
	}
	delete(s.driverStore, id)
	return &pb.DeleteDriverResponse{Success: true}, nil
}

func (s *server) ListDrivers(ctx context.Context, req *pb.ListDriversRequest) (*pb.ListDriversResponse, error) {
	var drivers []*pb.Driver
	for _, driver := range s.driverStore {
		drivers = append(drivers, driver)
	}
	return &pb.ListDriversResponse{Drivers: drivers}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDriverServiceServer(s, NewServer())
	log.Println("Server started at :80")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
