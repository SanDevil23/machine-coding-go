package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	snakes := []Snake{
		{Head: 14, Tail: 7},
		{Head: 31, Tail: 20},
		{Head: 38, Tail: 5},
		{Head: 84, Tail: 28},
		{Head: 97, Tail: 78},
	}

	ladders := []Ladder{
		{Start: 3, End: 22},
		{Start: 5, End: 8},
		{Start: 11, End: 26},
		{Start: 20, End: 29},
		{Start: 27, End: 56},
	}

	board := NewBoard(100, snakes, ladders)

	players := []*Player{
		{Name: "Alice", Position: 0},
		{Name: "Bob", Position: 0},
	}

	dice := &Dice{Sides: 6}

	game := &Game{
		Players: players,
		Board:   board,
		Dice:    dice,
	}

	game.Play()
}