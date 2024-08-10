package board

import (
	"errors"
	"fmt"
)

type Board struct {
	Size  int
	Board [][]string
}

func (b *Board) Init() {
	b.Size = 3
	b.Board = [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}

func (b *Board) Print() {
	for i := 0; i < b.Size; i++ {
		fmt.Println(b.Board[i])
	}
}

func (b *Board) SetCell(x int, y int, v string) error {
	value := &b.Board[x][y]
	if *value == "_" {
		*value = v
	} else {
		return errors.New("Invalid values")
	}
	return nil
}

func (b *Board) GetCell(x int, y int) string {
	return b.Board[x][y]
}

func (b *Board) CheckWinnerIs(s string) bool {
	// check lines
	counter, target, length := 0, 3, b.Size
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if b.Board[i][j] == s {
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
			if b.Board[j][i] == s {
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
		if b.Board[i][i] == s {
			counter++
		}
		if counter == target {
			return true
		}
	}
	// check diagonals
	counter = 0
	for i := 0; i < length; i++ {
		if b.Board[i][length-i-1] == s {
			counter++
		}
		if counter == target {
			return true
		}
	}
	return false
}
