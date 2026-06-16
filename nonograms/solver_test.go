package main

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSolveFiveByFiveDiamond(t *testing.T) {
	rowClues := [][]int{{1}, {3}, {5}, {3}, {1}}
	colClues := [][]int{{1}, {3}, {5}, {3}, {1}}

	got, err := Solve(rowClues, colClues)
	if err != nil {
		t.Fatalf("Solve returned error: %v", err)
	}

	want := strings.Join([]string{
		"..#..",
		".###.",
		"#####",
		".###.",
		"..#..",
	}, "\n")
	if got.GridString() != want {
		t.Fatalf("unexpected grid:\n%s\nwant:\n%s", got.GridString(), want)
	}
}

func TestSolutionStringIncludesPuzzleNumbers(t *testing.T) {
	rowClues := [][]int{{1}, {3}, {5}, {3}, {1}}
	colClues := [][]int{{1}, {3}, {5}, {3}, {1}}

	got, err := Solve(rowClues, colClues)
	if err != nil {
		t.Fatalf("Solve returned error: %v", err)
	}

	want := strings.Join([]string{
		"  1 3 5 3 1",
		"1 . . # . .",
		"3 . # # # .",
		"5 # # # # #",
		"3 . # # # .",
		"1 . . # . .",
	}, "\n")
	if got.String() != want {
		t.Fatalf("unexpected puzzle string:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestSolveAcceptsZeroAsEmptyClue(t *testing.T) {
	got, err := Solve([][]int{{0}, {1}, {0}}, [][]int{{1}})
	if err != nil {
		t.Fatalf("Solve returned error: %v", err)
	}

	if got.GridString() != ".\n#\n." {
		t.Fatalf("unexpected grid:\n%s", got.GridString())
	}
}

func TestSolveGeneratedTwentyFiveByTwentyFive(t *testing.T) {
	const size = 25
	grid := make([][]bool, size)
	for r := 0; r < size; r++ {
		grid[r] = make([]bool, size)
		for c := 0; c < size; c++ {
			v := (r*37 + c*17 + r*c*13 + r*r*5 + c*c*7) % 100
			grid[r][c] = v < 46
		}
	}

	rowClues, colClues := cluesFromGrid(grid)
	got, err := Solve(rowClues, colClues)
	if err != nil {
		t.Fatalf("Solve returned error: %v", err)
	}

	gotRows, gotCols := cluesFromGrid(got.BoolGrid())
	if !reflect.DeepEqual(gotRows, rowClues) {
		t.Fatalf("solution row clues do not match\n got: %v\nwant: %v", gotRows, rowClues)
	}
	if !reflect.DeepEqual(gotCols, colClues) {
		t.Fatalf("solution column clues do not match\n got: %v\nwant: %v", gotCols, colClues)
	}
}

func TestSolveTwentyFiveByThirtyPuzzle(t *testing.T) {
	const rowCluesEncoded = "4.2.7.2/4.2.3.4.2/3.4.3.6.1/1.5.3.7.2/3.2.7.5/2.2.4.5.3/2.3.6.1/1.2.1.6.3/3.1.3.2/6.8.2.1/6.1.9.1/2.1.10.3.2/3.12.1.2/2.10.4/3.5.4.6/3.5.5/3.2.6/1.1.6/5.3.3/8.3/1.10.4/14.1.6/15.7/8.4.3/2.1.1.3.2.3"
	const colCluesEncoded = "4.9.1/3.9.2/3.3.1.3.2/2.3.3/4.2.5/5.1.2.3/5.6/1.2.2.7/2.1.6/9.6/3.6.5/7.5/6.8.5/6.7.6/7.7.5/1.3.4.1.5/2.1.1.8.2/2.1.9/5.3.5.1/3.2.3/6.1/6.1/6.1.3.1/6.7.5.1/3.1.10/12/5.4/3.5.4/2.1.1.1.2.1.3/4.1.1.2.3"

	rowClues := parseEncodedClues(t, rowCluesEncoded)
	colClues := parseEncodedClues(t, colCluesEncoded)
	if len(rowClues) != 25 {
		t.Fatalf("decoded %d row clues, want 25", len(rowClues))
	}
	if len(colClues) != 30 {
		t.Fatalf("decoded %d column clues, want 30", len(colClues))
	}

	got, err := Solve(rowClues, colClues)
	if err != nil {
		t.Fatalf("Solve returned error: %v", err)
	}

	gotRows, gotCols := cluesFromGrid(got.BoolGrid())
	if !reflect.DeepEqual(gotRows, rowClues) {
		t.Fatalf("solution row clues do not match\n got: %v\nwant: %v", gotRows, rowClues)
	}
	if !reflect.DeepEqual(gotCols, colClues) {
		t.Fatalf("solution column clues do not match\n got: %v\nwant: %v", gotCols, colClues)
	}
}

func TestSolveRejectsImpossibleTotal(t *testing.T) {
	_, err := Solve([][]int{{2}}, [][]int{{1}, {0}})
	if err == nil {
		t.Fatal("expected error")
	}
	if errors.Is(err, ErrNoSolution) {
		t.Fatalf("expected validation error before search, got %v", err)
	}
}

func parseEncodedClues(t *testing.T, encoded string) [][]int {
	t.Helper()

	lines := strings.Split(encoded, "/")
	clues := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ".")
		clues[i] = make([]int, len(parts))
		for j, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				t.Fatalf("parse clue %q at line %d part %d: %v", part, i, j, err)
			}
			clues[i][j] = n
		}
	}
	return clues
}

func cluesFromGrid(grid [][]bool) (rows, cols [][]int) {
	rows = make([][]int, len(grid))
	for r := range grid {
		rows[r] = clueFromLine(grid[r])
	}

	if len(grid) == 0 {
		return rows, nil
	}

	cols = make([][]int, len(grid[0]))
	for c := range grid[0] {
		line := make([]bool, len(grid))
		for r := range grid {
			line[r] = grid[r][c]
		}
		cols[c] = clueFromLine(line)
	}

	return rows, cols
}

func clueFromLine(line []bool) []int {
	var clue []int
	run := 0
	for _, filled := range line {
		if filled {
			run++
			continue
		}
		if run > 0 {
			clue = append(clue, run)
			run = 0
		}
	}
	if run > 0 {
		clue = append(clue, run)
	}
	return clue
}
