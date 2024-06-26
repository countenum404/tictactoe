package main

import (
	"fmt"
	"shabashov/tic-tac-toe/internal/board"
)

/*
	Plans:
	1. To add configuration with marks
	2. Network version
	3. Player selection
*/

func main() {
	fmt.Println("Tic-Tac-Toe!")
	// select a player
	player1 := "X"
	player2 := "O"
	board := board.Board{}
	board.Init()
	var x1, y1, x2, y2 int
	for {
		board.Print()
		fmt.Println("Try", player1)
		fmt.Scan(&x1, &y1)
		board.SetCell(x1, y1, player1)

		if board.CheckWinnerIs(player1) {
			println("player1 is WINNER!!!")
			break
		}

		board.Print()
		fmt.Println("Try", player2)
		fmt.Scan(&x2, &y2)
		board.SetCell(x2, y2, player2)

		if board.CheckWinnerIs(player2) {
			println("player2 is WINNER!!!")
			break
		}
	}
}
