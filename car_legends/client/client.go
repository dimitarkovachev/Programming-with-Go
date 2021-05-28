package main

import (
	"context"
	"errors"
	"log"
	"time"

	grpcGeneratedCode "car_legends/grpc"

	"car_legends/client/resources"

	"google.golang.org/grpc"
)

var carsToAddSlice []*grpcGeneratedCode.Car

var carIndex *int = new(int)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}

	defer conn.Close()

	carsToAddSlice = resources.InitCarsToAddSlice()

	client := grpcGeneratedCode.NewCarLegendsClient(conn)

	runAddCar(client)
	runAddCars(client)
	runGetCars(client)
}

func getSingleCar(cars []*grpcGeneratedCode.Car) (*grpcGeneratedCode.Car, error) {
	if len(cars) > *carIndex {
		desiredCar := cars[*carIndex]

		*carIndex = *carIndex + 1

		return desiredCar, nil
	}

	return nil, errors.New("there are no more cars to send to the server")
}

func runAddCar(client grpcGeneratedCode.CarLegendsClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	carToAddViaAddCarMethod, err := getSingleCar(carsToAddSlice)

	if err != nil {
		log.Fatalf("Can't get car: %v", err)
	}

	_, err = client.AddCar(ctx, carToAddViaAddCarMethod)

	if err != nil {
		log.Fatalf("Could not add new car: %v", err)
	}
}

func runGetCars(client grpcGeneratedCode.CarLegendsClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.GetCars(ctx, &grpcGeneratedCode.Empty{})

	if err != nil {
		log.Fatalf("Could not get cars: %v", err)
	}

	log.Print(res)
}

func runAddCars(client grpcGeneratedCode.CarLegendsClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	stream, err := client.AddCars(ctx)

	if err != nil {
		log.Fatalf("Could not add cars: %v", err)
	}

	for i := 0; i < 5; i++ {
		carToSend, err := getSingleCar(carsToAddSlice)

		if err != nil {
			log.Fatalf("Can't get car: %v", err)
		}

		err = stream.Send(carToSend)

		if err != nil {
			log.Fatalf("Can't send car to server: %v %v %v", stream, carToSend, err)
		}
	}

	_, err = stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Can't close stream! %v", err)
	}
}
