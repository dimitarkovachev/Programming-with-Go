package resources

import (
	grpcGeneratedCode "github.com/dimitarkovachev/Programming-with-Go/car_legends/grpc"
)

func InitCarsToAddSlice() []*grpcGeneratedCode.Car {
	carsSlice := []*grpcGeneratedCode.Car{}

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "Lancia",
		Model:        "Delta",
		Year:         "1981",
		Notes:        "Integrale",
	})

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "Porsche",
		Model:        "911",
		Year:         "1991",
		Notes:        "Carrera",
	})

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "VW",
		Model:        "Golf",
		Year:         "1988",
		Notes:        "G60",
	})

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "Dodge",
		Model:        "Charger",
		Year:         "1970",
		Notes:        "Quarter mile at a time",
	})

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "BMW",
		Model:        "M3",
		Year:         "1986",
		Notes:        "E30",
	})

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "BMW",
		Model:        "M5",
		Year:         "2006",
		Notes:        "V10",
	})

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "Ferrarri",
		Model:        "F40",
		Year:         "1981",
		Notes:        "V12",
	})

	carsSlice = append(carsSlice, &grpcGeneratedCode.Car{
		Manufacturer: "Audi",
		Model:        "Quattro",
		Year:         "1981",
		Notes:        "Quattro!",
	})

	return carsSlice
}
