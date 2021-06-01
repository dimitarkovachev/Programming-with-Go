package main

import (
	"context"
	"io"
	"log"
	"net"
	strconv "strconv"

	grpcGeneratedCode "github.com/dimitarkovachev/Programming-with-Go/car_legends/grpc"

	sorter "github.com/dimitarkovachev/Programming-with-Go/sorter"
	"google.golang.org/grpc"
)

type carLegendsServer struct {
	grpcGeneratedCode.UnimplementedCarLegendsServer

	cars []*grpcGeneratedCode.Car
}

func (s *carLegendsServer) AddCar(ctx context.Context, car *grpcGeneratedCode.Car) (*grpcGeneratedCode.Empty, error) {
	s.cars = append(s.cars, car)

	log.Printf("Added car: %v", car)

	var e grpcGeneratedCode.Empty

	return &e, nil
}

func (s *carLegendsServer) GetCars(ctx context.Context, e *grpcGeneratedCode.Empty) (*grpcGeneratedCode.CarList, error) {
	s.SortCars()

	var carList = grpcGeneratedCode.CarList{
		Cars: s.cars,
	}

	return &carList, nil
}

func (s *carLegendsServer) AddCars(stream grpcGeneratedCode.CarLegends_AddCarsServer) error {
	for {
		car, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&grpcGeneratedCode.Empty{})
		}

		if err != nil {
			return err
		}

		s.cars = append(s.cars, car)

		log.Printf("Added car: %v", car)
	}
}

func (s *carLegendsServer) SortCars() {
	yearIndex := map[int]int{}
	years := []int{}
	sortedCars := []*grpcGeneratedCode.Car{}

	for i, v := range s.cars {
		carYearIntParsed, err := strconv.Atoi(v.Year)

		if err != nil {
			log.Printf("can't sort cars: %v\n", err)
			return
		}
		yearIndex[carYearIntParsed] = i
		years = append(years, carYearIntParsed)
	}

	sorter.SortSliceEventuallyBetter(years)

	for _, v := range years {
		sortedCars = append(sortedCars, s.cars[yearIndex[v]])
	}

	s.cars = sortedCars
}

func newServer() *carLegendsServer {
	s := &carLegendsServer{cars: initCarsSlice()}

	return s
}

func initCarsSlice() []*grpcGeneratedCode.Car {
	m5 := grpcGeneratedCode.Car{
		Manufacturer: "BMW",
		Model:        "M5",
		Year:         "1979",
		Notes:        "The first M5 and first luxury sports sedan.",
	}

	w124 := grpcGeneratedCode.Car{
		Manufacturer: "Mercedes",
		Model:        "W124",
		Year:         "1987",
		Notes:        "Together with W190 almost bankrupted Daimler.",
	}

	cars := []*grpcGeneratedCode.Car{}

	cars = append(cars, &m5)
	cars = append(cars, &w124)

	return cars
}

func main() {
	lis, err := net.Listen("tcp", ":50000")

	if err != nil {
		log.Fatalf("Can't listen: %v", err)
	}

	s := grpc.NewServer()
	grpcGeneratedCode.RegisterCarLegendsServer(s, newServer())
	s.Serve(lis)
}
