package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/himanshuvaish/golang-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	var (
		addDriverName     string
		addDriverWins     int
		addDriverPoles    int
		driverID          int
		updateDriverWins  int
		updateDriverPoles int
	)

	flag.StringVar(&addDriverName, "addName", "", "Name of the driver to add")
	flag.IntVar(&addDriverWins, "addWins", 0, "Number of wins for the driver to add")
	flag.IntVar(&addDriverPoles, "addPoles", 0, "Number of poles for the driver to add")
	flag.IntVar(&driverID, "driverID", 0, "ID of the driver for get, update, or delete operations")
	flag.IntVar(&updateDriverWins, "updateWins", 0, "Number of wins to update for the driver")
	flag.IntVar(&updateDriverPoles, "updatePoles", 0, "Number of poles to update for the driver")

	flag.Parse()

	conn, err := grpc.Dial("grpc.ecloudworx.com:80", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDriverServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Increased timeout
	defer cancel()

	if addDriverName != "" {
		addRes, err := c.AddDriver(ctx, &pb.AddDriverRequest{
			Driver: &pb.Driver{
				Name:  addDriverName,
				Wins:  int32(addDriverWins),
				Poles: int32(addDriverPoles),
			},
		})
		if err != nil {
			log.Fatalf("could not add driver: %v", err)
		}
		log.Printf("Driver added: %v", addRes.Driver)
	}

	if driverID > 0 {
		getRes, err := c.GetDriver(ctx, &pb.GetDriverRequest{Id: int32(driverID)})
		if err != nil {
			log.Fatalf("could not get driver: %v", err)
		}
		log.Printf("Driver retrieved: %v", getRes.Driver)
	}

	if driverID > 0 && (updateDriverWins > 0 || updateDriverPoles > 0) {
		updateRes, err := c.UpdateDriver(ctx, &pb.UpdateDriverRequest{
			Id:    int32(driverID),
			Wins:  int32(updateDriverWins),
			Poles: int32(updateDriverPoles),
		})
		if err != nil {
			log.Fatalf("could not update driver: %v", err)
		}
		log.Printf("Driver updated: %v", updateRes.Driver)
	}

	if driverID > 0 {
		deleteRes, err := c.DeleteDriver(ctx, &pb.DeleteDriverRequest{Id: int32(driverID)})
		if err != nil {
			log.Fatalf("could not delete driver: %v", err)
		}
		log.Printf("Driver deleted: %v", deleteRes.Success)
	}

	listRes, err := c.ListDrivers(ctx, &pb.ListDriversRequest{})
	if err != nil {
		log.Fatalf("could not list drivers: %v", err)
	}
	for _, driver := range listRes.Drivers {
		fmt.Printf("Driver: %v\n", driver)
	}
}
