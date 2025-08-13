package main

type VehicleType int;

const(
	none VehicleType = iota
	Car
	Bike
)

var VehicleTypeToString = map[VehicleType] string {
	none:		"none",
	Car:		"car",
	Bike:		"bike",
}

