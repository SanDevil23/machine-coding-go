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

// this function is only for the purpose of definition 
// we need to refine this function
func (p *Parking) createSlots() *Parking {
	slots := 10

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

	FirstFloor := Floor{
		Slots: totalSlots,
	}

	parking1 := Parking{
		Floors: []*Floor{&FirstFloor},
	}

	return &parking1
}