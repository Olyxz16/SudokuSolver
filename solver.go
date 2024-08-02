package main

import "fmt"

func Solve(sudoku *Sudoku) *Sudoku {
    step(sudoku)
	return sudoku
}

func step(sudoku *Sudoku) bool {
	if sudoku.remaining == 0 {
		return true
	}
	x, y, legals := sudoku.GetSmallestLegalSquare()
    if len(legals) == 0 {
        fmt.Println(sudoku.legalMovesMask)
		return false
	}
	for _, v := range legals {
		err := sudoku.Set(v, x, y)
		if err != nil {
			continue
		}
		ok := step(sudoku)
		if !ok {
			sudoku.Remove(x, y)
			continue
		}
        return true
	}
	return false
}
