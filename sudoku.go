package main

import (
	"fmt"
	"slices"
	"strings"
)

type Sudoku struct {
    gridSize int
    grid [][]int
    legalMoveGrid [][][]int
}

func New(gridSize int) Sudoku {
    grid := make([][]int, gridSize)
    for i:=0 ; i < gridSize ; i++ {
        grid[i] = make([]int, gridSize)
    }
    legalMoveGrid := make([][][]int, gridSize)
    return Sudoku{
        gridSize: gridSize,
        grid: grid, 
        legalMoveGrid: legalMoveGrid,
    }
}
func FromGrid(grid [][]int) Sudoku {
    sudoku := Sudoku{
        gridSize: len(grid)/3,
        grid: grid,
        legalMoveGrid: computeLegalMoves(grid),
    }
    return sudoku
}

func (sudoku Sudoku) Get(x,y int) int {
    return sudoku.grid[x][y]
}
func (sudoku Sudoku) Set(val,x,y int) {
    sudoku.grid[x][y] = val
}



func computeLegalMoves(grid [][]int) [][][]int {
    gridSize := len(grid)
    result := make([][][]int, 0)
    for i := 0 ; i < gridSize ; i++ {
        for j := 0 ; j < gridSize ; j++ {
            legalMovesAtPos := computeLegalMovesAtPos(grid, i,j)
            result[i][j] = legalMovesAtPos
        }
    }
    return result
}
func computeLegalMovesAtPos(grid [][]int, x, y int) []int {
    legals := []int{1,2,3,4,5,6,7,8,9}
    gridSize := len(grid)
    if grid[x][y] != 0 {
        return make([]int, 0)
    }
    
    // Compute row
    for i := 0 ; i < gridSize ; i++ {
        val := grid[x][i]
        if val != 0 {
            legals = remove(legals, val) 
        } 
    }
    
    // Compute column
    for i := 0 ; i < gridSize ; i++ {
        val := grid[i][y]
        if val != 0 {
            legals = remove(legals, val)
        }
    }

    return legals
}



func remove(arr []int, val int) []int {
    i := slices.Index(arr, val)
    if i == -1 {
        return arr
    }
    arr[i] = arr[len(arr)-1]
    return arr[:len(arr)-1]
}

func (sudoku Sudoku) Print() {
    var b strings.Builder
    for i:=0 ; i < sudoku.gridSize ; i++ {
        for j:=0 ; j < sudoku.gridSize ; j++ {
            fmt.Fprintf(&b, "%d ", sudoku.Get(i,j))
        }
        fmt.Fprintf(&b, "\n")
    }
}
