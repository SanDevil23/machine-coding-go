package main

type Slot struct {
	Car		bool
	Bike	bool
}

type Floor struct{
	Slots	[]*Slot
}

type Parking struct {
	Floors	[]*Floor
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


// this function is only for the purpose of definition 
// we need to refine this function
func (p *Parking) addSlots(floor *Floor, slots int) *Floor {

	totalSlots := []*Slot{}

	// adding slots to the floors
	for i:= slots; i>=0; i-- {
		if (i>slots/2){
			var s = Slot{
				Car: true,
				Bike: false,
			}
			totalSlots = append(totalSlots, &s)

		}else{
			var s = Slot{
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