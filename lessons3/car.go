package main

import "fmt"

type Car struct {
	Name              string
	EngineVolume      float64
	HorsePower        int
	YearOfManufacture int
	Fuel              string
	Transmission      string
}

func main() {

	CarOne := Car{
		Name:              "BMW 520",
		EngineVolume:      2.0,
		HorsePower:        184,
		YearOfManufacture: 2016,
		Fuel:              "Gasoline",
		Transmission:      "automatic",
	}

	CarTwo := Car{
		Name:              "BMW 320",
		EngineVolume:      2.2,
		HorsePower:        177,
		YearOfManufacture: 2009,
		Fuel:              "Diesel",
		Transmission:      "automatic",
	}
	// 1 задание
	fmt.Printf("%+v\n", CarOne)
	fmt.Printf("%+v\n", CarTwo)

	// 2 задание
	if CarOne.HorsePower > CarTwo.HorsePower {
		fmt.Printf("У %v больше лошадиных сил, чем у %v", CarOne.Name, CarTwo.Name)
	} else {
		fmt.Printf("У %v больше лошадиных сил, чем у %v", CarTwo.Name, CarOne.Name)
	}
	if CarOne.HorsePower == CarTwo.HorsePower {
		fmt.Printf("У %v столько же лошадиных сил, как и у %v", CarOne.Name, CarTwo.Name)
	}

}
