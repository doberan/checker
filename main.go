package main

import "fmt"

type Game struct {
	tern int,
	firstPlayer int,
	board Board,
}

type Board struct {
	cell [][]Cell
}

type State int

const CellState {
	NONE State = iota
	PLAYER1
	PLAYER2
}

func (s State) String() string {
	switch s {
	case NONE:
		return " "
	case PLAYER1:
		return "o"
	case PLAYER2:
		return "x"
	default:
		return " "
	}
}

type Cell struct {
	status CellState,
	// null or Stone
	stone Stone
}

func main() {
	var cell [][]int = [][]int{
		{2, 0, 2, 0, 2, 0, 2, 0},
		{0, 2, 0, 2, 0, 2, 0, 2},
		{2, 0, 2, 0, 2, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
	}
	showBoard(cell)
}
func showBoard(cell [][]int) {
	fmt.Println("   A   B   C   D   E   F   G   H  ")
	for x, xCell := range cell {
		fmt.Println(" + - + - + - + - + - + - + - + - +")
		fmt.Print(x)
		fmt.Print("| ")
		for _, xyCell := range xCell {
			fmt.Print(xyCell)
			fmt.Print(" | ")
		}
		fmt.Println()
	}
	fmt.Println(" + - + - + - + - + - + - + - + - +")
}
