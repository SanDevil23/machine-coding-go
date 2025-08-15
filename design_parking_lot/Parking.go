package main

type Slot struct {
	slotId	int64
	Car		bool
	Bike	bool
}

type Floor struct{
	Slots	[]*Slot
}

type Parking struct {
	ParkingId string
	Floors    []*Floor
}

func NewParking(parkingId string) *Parking {
	return &Parking{
		ParkingId: parkingId,
		Floors:    []*Floor{},
	}
}

func (p *Parking) addFloor(floor *Floor) *Parking {
	p.Floors = append(p.Floors, floor)
	return p
}

func (p *Parking) NewFloor(slots int) *Parking{
	floor := &Floor{
		Slots: []*Slot{},
	}
	p.addSlots(floor, slots)
	p.addFloor(floor)
	return p
}

// TODO:  currently dividing slots into 50-50 for each vehicle type
// custom logic needed for slot allotment
func (p *Parking) addSlots(floor *Floor, slots int) *Floor {

	totalSlots := []*Slot{}

	// adding slots to the floors
	for i:= slots; i>=0; i-- {
		if (i>slots/2){
			var s = Slot{
				slotId: int64(i),
				Car: true,
				Bike: false,
			}
			totalSlots = append(totalSlots, &s)

		}else{
			var s = Slot{
				slotId: int64(i),
				Car: false,
				Bike: true,
			}	
			totalSlots = append(totalSlots, &s)
		}
	}

	floor = &Floor{
		Slots: totalSlots,
	}

	return floor
}