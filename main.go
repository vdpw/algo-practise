package main

import (
	"algo/nonograms"
	"fmt"
	"time"
)

func printLine(item []bool) {
	for _, b := range item {
		if b {
			fmt.Print("⏹️ ")
		} else {
			fmt.Print("▶️ ")
		}
	}
	fmt.Println()
}

func main() {
	leftC := nonograms.Constraints{
		nonograms.Constraint{3, 8, 4, 2, 3, 1},
		nonograms.Constraint{1, 3, 3, 3, 3, 3, 4, 1, 1},
		nonograms.Constraint{1, 3, 1, 4, 1, 5, 1},
		nonograms.Constraint{2, 6, 1, 5, 2, 5, 2},
		nonograms.Constraint{1, 6, 6, 1, 1},
		nonograms.Constraint{4, 1, 1, 1, 4, 3},
		nonograms.Constraint{2, 3, 2, 1, 4, 2, 1, 1, 4},
		nonograms.Constraint{4, 10, 1, 6, 7},
		nonograms.Constraint{3, 3, 1, 9, 5, 5},
		nonograms.Constraint{3, 4, 4, 1, 1, 4, 3},
		nonograms.Constraint{1, 1, 3, 3, 5, 6, 1, 1, 1, 1, 1},
		nonograms.Constraint{10, 2, 10, 2, 5, 1},
		nonograms.Constraint{12, 3, 10, 3, 7},
		nonograms.Constraint{1, 7, 1, 7, 2, 4, 4},
		nonograms.Constraint{1, 1, 1, 1, 4, 7, 4, 1, 5},
		nonograms.Constraint{1, 3, 1, 1, 2, 7},
		nonograms.Constraint{9, 4, 4},
		nonograms.Constraint{1, 2, 3, 8, 4, 2},
		nonograms.Constraint{5, 3, 7, 1, 3, 4, 1},
		nonograms.Constraint{5, 1, 6, 1, 5, 1, 1},
		nonograms.Constraint{4, 1, 1, 4, 2, 3, 3, 3},
		nonograms.Constraint{4, 1, 1, 6, 1, 2, 3, 3},
		nonograms.Constraint{4, 1, 1, 1, 4, 3, 1, 1, 4},
		nonograms.Constraint{5, 4, 7, 3, 1, 1, 2, 1},
		nonograms.Constraint{5, 4, 7, 3, 2, 1, 4},
		nonograms.Constraint{5, 4, 7, 2, 1, 11, 2, 1, 1},
		nonograms.Constraint{7, 4, 6, 1, 6, 2, 6},
		nonograms.Constraint{6, 5, 8, 9, 3},
		nonograms.Constraint{5, 2, 2, 1, 3, 1, 4, 2, 4},
		nonograms.Constraint{6, 2, 8, 1, 2, 2, 3},
		nonograms.Constraint{2, 3, 1, 1, 2, 4, 4, 3, 3, 3},
		nonograms.Constraint{2, 1, 2, 6, 4, 6, 3, 2},
		nonograms.Constraint{3, 1, 7, 3, 6, 1, 2},
		nonograms.Constraint{2, 5, 3, 5, 2, 1, 5, 2},
		nonograms.Constraint{8, 3, 5, 3, 3, 2, 2},
		nonograms.Constraint{3, 5, 1, 4, 3, 6, 3, 1, 2},
		nonograms.Constraint{3, 3, 21, 3, 3, 2},
		nonograms.Constraint{1, 1, 1, 9, 1, 5, 3, 2, 1},
		nonograms.Constraint{1, 2, 10, 8, 3, 3},
		nonograms.Constraint{2, 1, 3, 7, 1, 6, 2, 2},
		nonograms.Constraint{2, 1, 4, 2, 2, 7, 1, 1, 2},
		nonograms.Constraint{3, 1, 1, 13, 2, 1, 5, 5},
		nonograms.Constraint{12, 3, 1, 1, 1, 3, 8},
		nonograms.Constraint{1, 1, 1, 4, 6, 6, 1, 2, 5, 4, 3},
		nonograms.Constraint{3, 5, 2, 1, 3, 3, 3, 4, 10},
		nonograms.Constraint{14, 7, 4, 1, 10},
		nonograms.Constraint{9, 5, 5, 3, 8},
		nonograms.Constraint{2, 10, 4, 3, 5, 4},
		nonograms.Constraint{1, 3, 3, 1, 3, 3, 2, 5},
		nonograms.Constraint{4, 1, 1, 1, 5, 5, 1, 5},
	}

	topC := nonograms.Constraints{
		nonograms.Constraint{4, 3, 3, 11, 3, 1, 1, 5, 6},
		nonograms.Constraint{1, 7, 10, 3, 1, 1, 3, 4},
		nonograms.Constraint{4, 2, 12, 2, 1, 4},
		nonograms.Constraint{1, 3, 15, 1, 2},
		nonograms.Constraint{1, 3, 3, 8, 3, 3, 4},
		nonograms.Constraint{5, 6, 1, 5, 5, 4},
		nonograms.Constraint{11, 1, 1, 4, 7},
		nonograms.Constraint{13, 2, 5, 8},
		nonograms.Constraint{2, 3, 3, 2, 4, 4, 9},
		nonograms.Constraint{5, 1, 4, 3, 7, 1, 4, 4, 1, 3},
		nonograms.Constraint{5, 3, 2, 2, 5, 3, 1, 5, 3, 1},
		nonograms.Constraint{3, 2, 1, 6, 1, 5, 4},
		nonograms.Constraint{2, 2, 3, 10},
		nonograms.Constraint{2, 1, 2, 9},
		nonograms.Constraint{4, 3, 1, 3, 2, 5, 2},
		nonograms.Constraint{2, 7, 9, 1, 7},
		nonograms.Constraint{2, 6, 3, 6, 7, 3},
		nonograms.Constraint{1, 3, 1, 12, 8, 3},
		nonograms.Constraint{1, 1, 1, 6, 9, 9, 1},
		nonograms.Constraint{2, 2, 3, 14, 16, 1},
		nonograms.Constraint{5, 1, 3, 21, 9},
		nonograms.Constraint{5, 4, 6, 3, 2, 3, 3, 2},
		nonograms.Constraint{3, 5, 6, 7, 3, 1},
		nonograms.Constraint{3, 5, 2, 5, 3, 1, 1},
		nonograms.Constraint{6, 4, 3, 4, 6, 1},
		nonograms.Constraint{2, 4, 4, 2, 2, 1, 6},
		nonograms.Constraint{5, 5, 3, 8},
		nonograms.Constraint{1, 1, 5, 3, 3, 2, 5},
		nonograms.Constraint{3, 5, 1, 3, 1, 1, 1, 7},
		nonograms.Constraint{2, 8, 1, 2, 3, 3, 1, 3},
		nonograms.Constraint{2, 4, 3, 3, 7, 2},
		nonograms.Constraint{5, 1, 3, 5},
		nonograms.Constraint{10, 2, 2, 1, 5, 3},
		nonograms.Constraint{4, 4, 2, 4, 1, 3, 6, 3},
		nonograms.Constraint{1, 3, 1, 10, 7, 5},
		nonograms.Constraint{1, 1, 9, 1, 4, 3, 1},
		nonograms.Constraint{1, 4, 3, 1, 5, 2},
		nonograms.Constraint{2, 3, 4, 5, 2},
		nonograms.Constraint{2, 1, 1, 1, 3, 8, 1, 3},
		nonograms.Constraint{1, 3, 1, 1, 1, 5, 1, 1},
		nonograms.Constraint{3, 3, 3, 5},
		nonograms.Constraint{3, 1, 4, 1, 5, 3, 7},
		nonograms.Constraint{4, 5, 1, 7, 5, 2, 8},
		nonograms.Constraint{4, 2, 4, 1, 3, 5, 4, 6},
		nonograms.Constraint{3, 1, 6, 1, 1, 3, 2, 4, 8},
		nonograms.Constraint{2, 2, 3, 3, 1, 3, 2, 6},
		nonograms.Constraint{1, 1, 4, 7, 3, 1, 1, 1, 11},
		nonograms.Constraint{26, 3, 8},
		nonograms.Constraint{1, 5, 2, 3, 11, 1, 4, 2},
		nonograms.Constraint{5, 5, 1, 6, 14, 1, 2, 2},
	}

	p := nonograms.NewPuzzle(50, 50, leftC, topC)
	now := time.Now()
	p.Resolve()
	fmt.Println(time.Since(now).Microseconds())
	for _, row := range p.Grids {
		printLine(row)
	}
}
