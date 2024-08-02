package main

import "fmt"

func main() {
	sudoku := New(40)
	sudoku.Print()
	fmt.Println(sudoku.remaining)

	solved := Solve(sudoku)
	solved.Print()
    fmt.Println(sudoku.remaining)
}
