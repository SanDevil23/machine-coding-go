package main

import "fmt"

type Snake struct {
	Head int
	Tail int
}

type Ladder struct {
	Start int
	End   int
}

type Board struct {
	Size    int
	Snakes  map[int]int
	Ladders map[int]int
}

// new board function to initialize the board
func newBoard(size int, snakes []Snake, ladders []Ladder) *Board {
	snakeMap := make(map[int]int)
	ladderMap := make(map[int]int)

	for _, s := range snakes {
		snakeMap[s.Head] = s.Tail
	}

	for _, l := range ladders {
		ladderMap[l.Start] = l.End
	}

	return &Board{
		Size:    size,
		Snakes:  snakeMap,
		Ladders: ladderMap,
	}
}

func (b *Board) adjustPosition(pos int) int {
	// check if the position has a snake
	if finalPos, ok := b.Snakes[pos]; ok {
		fmt.Println("Bitten by Snake")
		return finalPos;
	}

	// check if the position has a ladder
	if finalPos,ok := b.Ladders[pos]; ok {
		fmt.Println("Climbing the ladder")
		return finalPos
	}

	return pos
}