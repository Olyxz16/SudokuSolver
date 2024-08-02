package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

type Sudoku struct {
    gridSize int
    remaining int
    grid []int
    legalMovesMask []int
}


func New(baseNumbers int) *Sudoku {
    sudoku := BuildFullValidGrid(3)
    cellCount := len(sudoku.grid)
    sudoku.remaining = cellCount - baseNumbers
    inds := rand.Perm(cellCount)
    for i:=0 ; i < sudoku.remaining ; i++ {
        ind := inds[i]
        sudoku.grid[ind] = 0 
    }
    sudoku.legalMovesMask = sudoku.computeLegalMovesMask()
    return sudoku
}
// Build valid grid : 
// create empty grid
// for each number n
// fill all column with number n where legal
func BuildFullValidGrid(gridSize int) *Sudoku {
    sudoku := Empty(gridSize) 
    for v:=1 ; v<=9 ; v++ {
        for x:=0 ; x<3*gridSize ; x++ {
            remainingY := rand.Perm(9)
            for _, y := range remainingY {
                ind := sudoku.indFromPos(x,y)
                if sudoku.isLegal(v, ind) {
                    sudoku.Set(v,x,y)
                    break
                }
            }
        } 
    }
    if sudoku.remaining > 1 {
        return BuildFullValidGrid(gridSize)
    }
    return sudoku
}
func Empty(gridSize int) *Sudoku {
    cellCount := 9*gridSize*gridSize
    grid := make([]int, cellCount)
    var legalMovesMask []int
    sudoku := &Sudoku{
        gridSize: gridSize,
        remaining: 9*gridSize*gridSize,
        grid: grid, 
        legalMovesMask: legalMovesMask,
    }
    sudoku.legalMovesMask = sudoku.computeLegalMovesMask()
    return sudoku
}




func (sudoku *Sudoku) Remaining() int {
    return sudoku.remaining    
}
func (sudoku *Sudoku) Get(x,y int) int {
    ind := sudoku.indFromPos(x,y)
    return sudoku.grid[ind]
}
func (sudoku *Sudoku) Set(val,x,y int) error {
    ind := sudoku.indFromPos(x,y)
    if !sudoku.isLegal(val, ind) {
        return fmt.Errorf("This move is illegal")
    }
    if sudoku.grid[ind] == 0 {
        sudoku.remaining--
    }
    sudoku.grid[ind] = val
    sudoku.legalMovesMask = sudoku.computeLegalMovesMask()
    return nil
}
func (sudoku *Sudoku) Remove(x,y int) {
    ind := sudoku.indFromPos(x,y)
    if sudoku.grid[ind] != 0 {
        sudoku.remaining++
    }
    sudoku.grid[ind] = 0
    sudoku.legalMovesMask = sudoku.computeLegalMovesMask()
}
func (sudoku *Sudoku) indFromPos(x,y int) (ind int) {
    return 3*sudoku.gridSize*x + y
}
func (sudoku *Sudoku) isLegal(val, ind int) bool {
    mask := sudoku.legalMovesMask[ind]
    return mask & (1 << val) > 0
}


func (sudoku *Sudoku) GetSmallestLegalSquare() (x,y int, mask []int) {
    var ind, vmin int
    vmin = 10
    for i, v := range sudoku.legalMovesMask {
        arr := maskToArray(v)
        count := len(arr)
        legit := (count < vmin && count > 0)
        if legit {
            ind = i
            vmin = count 
            mask = arr
        }
    }
    x = ind / (3*sudoku.gridSize)
    y = ind % (3*sudoku.gridSize)
    return x,y,mask
}
func maskToArray(mask int) []int {
    var result []int
    var step int
    for mask > 0 {
        if int(mask & 1) == 1 {
            result = append(result, step) 
        }
        step++
        mask >>= 1
    }
    return result
}
func (sudoku *Sudoku) computeLegalMovesMask() []int {
    cellCount := len(sudoku.grid)
    result := make([]int, cellCount)
    for i := 0 ; i < 3*sudoku.gridSize ; i++ {
        for j := 0 ; j < 3*sudoku.gridSize ; j++ {
            mask := sudoku.computeLegalMovesAt(i,j)
            result[sudoku.indFromPos(i,j)] = mask
        }
    }
    return result
}
func (sudoku *Sudoku) computeLegalMovesAt(x, y int) int {
    legals := []int{ 1,2,3,4,5,6,7,8,9 }
    gridSize := 3*sudoku.gridSize
    if sudoku.Get(x,y) != 0 {
        return 0
    }
    
    // Compute row
    for i := 0 ; i < gridSize ; i++ {
        val := sudoku.Get(x, i)
        if val != 0 {
            legals = remove(legals, val) 
        } 
    }
    
    // Compute column
    for i := 0 ; i < gridSize ; i++ {
        val := sudoku.Get(i, y)
        if val != 0 {
            legals = remove(legals, val)
        }
    }

    // Compute square
    sqIndX := x / 3
    sqIndY := y / 3
    for i := 3*sqIndX ; i < 3*sqIndX+3 ; i++ {
        for j := 3*sqIndY ; j < 3*sqIndY+3 ; j++ {
            val := sudoku.Get(i,j)
            if val != 0 {
                legals = remove(legals, val)
            }
        }
    }
    
    var mask int
    for _, val := range legals {
        mask |= 1 << val 
    }
    return mask
}


func remove(arr []int, val int) []int {
    i := slices.Index(arr, val)
    if i == -1 {
        return arr
    }
    arr[i] = arr[len(arr)-1]
    return arr[:len(arr)-1]
}

func (sudoku *Sudoku) Print() {
    var b strings.Builder
    for i:=0 ; i < 3*sudoku.gridSize ; i++ {
        if i % 3 == 0 {
            for j:=0 ; j < 3*(sudoku.gridSize+1)+1 ; j++ {
               fmt.Fprintf(&b, "-") 
            }
            fmt.Fprintf(&b, "\n")
        }
        for j:=0 ; j < 3*sudoku.gridSize ; j++ {
            if j % 3 == 0{
                fmt.Fprintf(&b, "|")
            }
            fmt.Fprintf(&b, "%d", sudoku.Get(i,j))
        }
        fmt.Fprintf(&b, "|\n")
    }
    for j:=0 ; j < 3*(sudoku.gridSize+1)+1 ; j++ {
        fmt.Fprintf(&b, "-") 
    }
    fmt.Fprintf(&b, "\n")

    fmt.Print(b.String())
}
