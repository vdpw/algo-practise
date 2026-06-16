package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const exampleColClues = "4.2.7.2/4.2.3.4.2/3.4.3.6.1/1.5.3.7.2/3.2.7.5/2.2.4.5.3/2.3.6.1/1.2.1.6.3/3.1.3.2/6.8.2.1/6.1.9.1/2.1.10.3.2/3.12.1.2/2.10.4/3.5.4.6/3.5.5/3.2.6/1.1.6/5.3.3/8.3/1.10.4/14.1.6/15.7/8.4.3/2.1.1.3.2.3"
const exampleRowClues = "4.9.1/3.9.2/3.3.1.3.2/2.3.3/4.2.5/5.1.2.3/5.6/1.2.2.7/2.1.6/9.6/3.6.5/7.5/6.8.5/6.7.6/7.7.5/1.3.4.1.5/2.1.1.8.2/2.1.9/5.3.5.1/3.2.3/6.1/6.1/6.1.3.1/6.7.5.1/3.1.10/12/5.4/3.5.4/2.1.1.1.2.1.3/4.1.1.2.3"

func main() {
	rowClues, err := parseClueSpec(exampleRowClues)
	if err != nil {
		log.Fatalf("parse row clues: %v", err)
	}

	colClues, err := parseClueSpec(exampleColClues)
	if err != nil {
		log.Fatalf("parse column clues: %v", err)
	}

	solution, err := Solve(rowClues, colClues)
	if err != nil {
		log.Fatalf("solve puzzle: %v", err)
	}

	fmt.Printf("Solved %dx%d puzzle:\n%s\n", solution.Width, solution.Height, solution.String())
}

func parseClueSpec(encoded string) ([][]int, error) {
	lines := strings.Split(encoded, "/")
	clues := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ".")
		clues[i] = make([]int, len(parts))
		for j, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("line %d part %d %q: %w", i, j, part, err)
			}
			clues[i][j] = n
		}
	}
	return clues, nil
}
