package main


type Board struct {
	size  int
	grid  [][]string
}


func NewBoard() *Board{
	size := 3 // Default size for a Tic Tac Toe board
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
	}
	return &Board{
		size: size,
		grid: grid,
	}	
}

func (b *Board) isCellEmpty(row int, col int) bool {
	if row < 0 || row >= b.size || col < 0 || col >= b.size {
		return false
	}
	return b.grid[row][col] == ""
}

func (b *Board) makeMove(row int, col int, mark string) bool{
	if empty:=b.isCellEmpty(row, col); !empty {
		return false
	}
	b.grid[row][col] = mark
	return true
}


