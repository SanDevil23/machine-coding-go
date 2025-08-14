package main

type ParkingSpaces struct {
	Parking []*Parking
}

// function for initializing parking space struct
func NewParkingSpaces() *ParkingSpaces {
	return &ParkingSpaces{
		Parking: []*Parking{},
	}
}

// function for adding a new parking space into the parking spaces
func (ps *ParkingSpaces) AddParking(parking *Parking) {
	ps.Parking = append(ps.Parking, parking);
}

