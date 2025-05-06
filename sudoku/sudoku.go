package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Panel [9][9]int
}

func NewBoard() *Board {
	return &Board{Panel: [9][9]int{}}
}

func (b *Board) findEmptyCell() (int, int) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b.Panel[r][c] == 0 {
				return r, c
			}
		}
	}
	return -1, -1
}

func (b *Board) check(row, column int, value int) bool {
	// checking through the row
	for i := 0; i < 9; i++ {
		if b.Panel[row][i] == value {
			return false
		}
	}

	// checking through the column
	for i := 0; i < 9; i++ {
		if b.Panel[i][column] == value {
			return false
		}
	}

	// checking through the 3x3 block of the cell
	secRow := row / 3
	secCol := column / 3
	for i := secRow * 3; i < (secRow*3 + 3); i++ {
		for j := (secCol * 3); j < (secCol*3 + 3); j++ {
			if b.Panel[i][j] == value {
				return false
			}
		}
	}

	return true
}

func (b *Board) Solve() bool {
	row, column := b.findEmptyCell()

	if row >= 0 && column >= 0 {
		for v := 1; v < 10; v++ {
			if b.check(row, column, v) {
				b.Panel[row][column] = v
				if b.Solve() {
					return true
				}
				// backtracking if the board cannot be solved using current configuration
				b.Panel[row][column] = 0
			}
		}
	} else {
		// if the board is complete
		return true
	}

	// returning false the board cannot be solved using current configuration
	return false
}

func (b *Board) print(arr []int, last bool) {
	strList := []string{}
	for i := range arr {
		strList = append(strList, strconv.Itoa(arr[i]))
	}
	s := strings.Join(strList, ", ")
	if last {
		fmt.Println(s)
	} else {
		fmt.Printf("%s | ", s)
	}
}

func (b *Board) Print() {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("- - - - - - - - - - - - - -")
		}

		b.print(b.Panel[i][0:3], false)
		b.print(b.Panel[i][3:6], false)
		b.print(b.Panel[i][6:9], true)
	}

}
