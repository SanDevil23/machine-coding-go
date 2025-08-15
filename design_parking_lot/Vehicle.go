package main

type Vehicle struct {
	VType 	*VehicleType
	RegNo	string
	color 	string
}

func NewVehicle(vType *VehicleType, regNo string, color string) *Vehicle {
	return &Vehicle{
		VType: vType,
		RegNo: regNo,
		color: color,
	}
}

func (v *Vehicle) GetType() *VehicleType {
	return v.VType
}

func (v *Vehicle) GetRegNo() string {
	return v.RegNo
}

func (v *Vehicle) GetColor() string {
	return v.color
}
