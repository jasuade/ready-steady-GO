package main

import (
	"fmt"
	"strings"
)

func main() {
	//Creation of a tic-tac-toe borad
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//Moking players taking turns
	board[0][2] = "X"
	board[1][0] = "O"
	board[2][1] = "X"
	board[0][1] = "O"
	board[2][2] = "X"
	board[1][1] = "O"
	board[2][0] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " ")) //to print the board 3 iterations over each string and join the inner stings
	}
}
