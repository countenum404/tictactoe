package tests

import (
	"shabashov/tic-tac-toe/internal/board"
	"testing"
)

func InitializedBoard() board.Board {
	b := board.Board{}
	b.Init()
	return b
}

func TestBoardInit(t *testing.T) {
	b := InitializedBoard()
	iVal := "_"
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.GetCell(i, j) != iVal {
				t.Errorf("Initial values are incorrect in x=%d y=%d", j, i)
			}
		}
	}
}

func TestSetCell(t *testing.T) {
	b := InitializedBoard()
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			b.SetCell(j, i, "X")
		}
	}
	// check board filled
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.GetCell(j, i) != "X" {
				t.Errorf("Value not setted in x=%d y=%d", j, i)
			}
		}
	}
	// check cannot overwrite cell
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			b.SetCell(j, i, "O")
			/*
				TODO:
				Change function to return the error instead of logging
			*/
		}
	}

}

func TestCheckIsWinnerRowsAndColumns(t *testing.T) {
	b := InitializedBoard()
	t.Logf(b.GetCell(0, 0))
	player := "X"
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			b.SetCell(i, j, player)
		}
		if b.CheckWinnerIs(player) != true {
			t.Errorf("Player is not winner at row %d", i)
			b.Print()
		}
		b.Init()
	}

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			b.SetCell(j, i, player)
		}
		if b.CheckWinnerIs(player) != true {
			t.Errorf("Player is not winner at column %d", i)
			b.Print()
		}
		b.Init()
	}
}

func TestCheckIsWinnerDiagonals(t *testing.T) {
	b := InitializedBoard()
	t.Logf(b.GetCell(0, 0))
	player := "X"
	b.SetCell(0, 0, player)
	b.SetCell(1, 1, player)
	b.SetCell(2, 2, player)
	if b.CheckWinnerIs(player) != true {
		t.Errorf("Player is not winner at Main diagonal")
		b.Print()
	}
	b.Init()

	b.SetCell(0, 2, player)
	b.SetCell(1, 1, player)
	b.SetCell(2, 0, player)
	if b.CheckWinnerIs(player) != true {
		t.Errorf("Player is not winner at Side diagonal")
		b.Print()
	}
	b.Init()
}
