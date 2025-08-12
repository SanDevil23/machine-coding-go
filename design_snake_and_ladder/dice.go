package main

import "math/rand"

type Dice struct {
	Sides int
}

func (d *Dice) Roll() int {
	return rand.Intn(d.Sides) + 1;
}