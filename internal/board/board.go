package board

import "fmt"

type Board struct {
	Size  int
	board [][]string
}

func (b *Board) Init() {
	b.Size = 3
	b.board = [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}

func (b *Board) Print() {
	for i := 0; i < b.Size; i++ {
		fmt.Println(b.board[i])
	}
}

func (b *Board) SetCell(x int, y int, v string) {
	value := &b.board[x][y]
	if *value == "_" {
		*value = v
	} else {
		fmt.Println(x, y, "values isn't valid")
	}
}

func (b *Board) GetCell(x int, y int) string {
	return b.board[x][y]
}

func (b *Board) CheckWinnerIs(s string) bool {
	// check lines
	counter, target, length := 0, 3, b.Size
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if b.board[i][j] == s {
				counter++
			}
		}
		if counter == target {
			return true
		}
		counter = 0
	}
	// check columns
	counter = 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if b.board[j][i] == s {
				counter++
			}
		}
		if counter == target {
			return true
		}
		counter = 0
	}

	// check diagonals +
	counter = 0
	for i := 0; i < length; i++ {
		if b.board[i][i] == s {
			counter++
		}
		if counter == target {
			return true
		}
	}
	// check diagonals
	counter = 0
	for i := 0; i < length; i++ {
		if b.board[i][length-i-1] == s {
			counter++
		}
		if counter == target {
			return true
		}
	}
	return false
}
