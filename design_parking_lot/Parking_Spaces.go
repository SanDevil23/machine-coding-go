package main

type ParkingSpace struct {
	Parking []*Parking
}

func (ps *ParkingSpace) AddParking(parking *Parking) {
	ps.Parking = append(ps.Parking, parking)
}

