package main

import "fmt"

type Ticket struct {
	TicketId    string
	VehicleType *VehicleType
	VehicleNo   string
	ParkingId   string
}

func NewTicket(ticketId string, vehicle *Vehicle, parkingId string) *Ticket {
	return &Ticket{
		TicketId:    ticketId,
		VehicleType: vehicle.VType,
		VehicleNo:   vehicle.RegNo,
		ParkingId:   parkingId,
	}
}

func (t *Ticket) generateTicketId(vt *VehicleType, park *Parking) string {
	slot, floor := t.findParking(vt, park)
	if slot != nil && floor != -1 {
		return fmt.Sprintf("%s_%d_%d", "PR1234", floor, slot.slotId)
	}
	return "Parking Unavailable"
}

func (t *Ticket) findParking(vt *VehicleType, park *Parking) (*Slot, int) {

	for indF, floor := range park.Floors {
		for _, slot := range floor.Slots {
			if slot.Car && vt.toString() == "car" {
				return slot, indF
			}
			if slot.Bike && vt.toString() == "bike" {
				return slot, indF
			}
			fmt.Println("No suitable parking found")
			return nil, -1
		}
	}

	fmt.Println("No suitable parking found")
	return nil, -1
}