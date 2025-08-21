package main

import "fmt"

type Game struct {
	Player1       Player
	Player2       Player
	Board         *Board
	CurrentPlayer *Player
}

func NewGame() *Game {
	p1, p2 := setupPlayers()
	board := NewBoard()
	currentPlayer := &p1

	return &Game{
		Player1:       p1,
		Player2:       p2,
		Board:         board,
		CurrentPlayer: currentPlayer,
	}
}

func setupPlayers() (Player, Player) {
	p1 := NewPlayer("Player 1", "X")
	p2 := NewPlayer("Player 2", "O")
	return *p1, *p2
}

func (g *Game) switchPlayer() {
	if g.CurrentPlayer == &g.Player1 {
		g.CurrentPlayer = &g.Player2
	} else {
		g.CurrentPlayer = &g.Player1
	}
}

func (g *Game) start() {
	for true {
		var row, col int

		for true {
			fmt.Print("Enter row and column (0-2): ")
			_, err := fmt.Scanf("%d %d", &row, &col)
			if err != nil {
				fmt.Println("Invalid input. Please enter numbers.")
				continue
			}
			if row < 0 || row > 2 || col < 0 || col > 2 {
				fmt.Println("Invalid position. Please enter values between 0 and 2.")
				continue
			}
			if !g.Board.makeMove(row, col, g.CurrentPlayer.GetSymbol()) {
				fmt.Println("Position already taken. Try again.")
				continue
			}
			break
		}

		if g.Board.checkWin(g.CurrentPlayer.GetSymbol()) {
			fmt.Printf("%s wins!\n", g.CurrentPlayer.GetName())
			return
		}else if g.Board.isFull() {
			fmt.Println("It's a draw!")
			return
		}else{
			g.switchPlayer()
			fmt.Println("Next player's turn.")
		}
	}
}
