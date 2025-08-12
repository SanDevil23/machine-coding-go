package main

import "fmt"

type Game struct {
	Board   *Board
	Dice    *Dice
	Players []*Player
}

func (g *Game) Play() {
	playerQueue := make([]*Player, len(g.Players))
	copy(playerQueue, g.Players)

	for {
		current := playerQueue[0]
		playerQueue = playerQueue[1:]

		fmt.Println("\n" + current.Name + "'s turn")
		roll := g.Dice.Roll()
		fmt.Println("Dice Rolled: ", roll)

		nextPos := current.Position + roll
		
		if nextPos > g.Board.Size {
			println("Can't move, stay in place.")
		}else {
			nextPos = g.Board.AdjustPosition(nextPos)
			current.Position = nextPos
			println(current.Name, "moved to:", current.Position)
		}

		// check win condition
		if current.Position == g.Board.Size {
			println("Congratulations! " + current.Name + " wins!")
			break
		}

		// adding the current player to the end of the queue
		playerQueue = append(playerQueue, current)
	}

}