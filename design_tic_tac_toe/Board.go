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

func (b *Board) isFull() bool {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.isCellEmpty(i, j) {
				return false
			}
		}
	}
	return true
}

func (b *Board) checkWin(mark string) bool {
	// check rows and columns
	for i := 0; i < b.size; i++ {
		if b.grid[i][0] == mark && b.grid[i][1] == mark && b.grid[i][2] == mark {
			return true
		}
		if b.grid[0][i] == mark && b.grid[1][i] == mark && b.grid[2][i] == mark {
			return true
		}
	}
	// check diagonals
	if b.grid[0][0] == mark && b.grid[1][1] == mark && b.grid[2][2] == mark {
		return true
	}
	if b.grid[0][2] == mark && b.grid[1][1] == mark && b.grid[2][0] == mark {
		return true
	}
	return false
}
