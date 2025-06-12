package nonograms

import "fmt"

type Constraint []int

type Constraints []Constraint

func (c Constraint) Total() int {
	sum := 0
	for _, n := range c {
		sum += n
	}
	return sum
}

type Puzzle struct {
	TopConstraints  Constraints
	LeftConstraints Constraints
	Grids           [][]bool
}

// across Constraint c to find all possible rows
func GetPossiblesRows(length int, c Constraint) [][]bool {
	parts := len(c)
	spacesToArrange := length - (parts - 1) - c.Total()

	plans := allArranges(spacesToArrange, parts)

	result := [][]bool{}
	for i := range plans {
		plan := plans[i]
		r := []bool{}
		for j := range plan {
			if j > 0 {
				plan[j] = plan[j] + 1
			}

			for k := 0; k < plan[j]; k++ {
				r = append(r, false)
			}
			for k := 0; k < c[j]; k++ {
				r = append(r, true)
			}
		}
		for len(r) < length {
			r = append(r, false)
		}

		result = append(result, r)
	}

	return result
}

func allArranges(total, parts int) [][]int {
	if parts == 0 {
		return nil
	}
	if parts == 1 {
		res := [][]int{}
		for i := 0; i <= total; i++ {
			res = append(res, []int{i})
		}
		return res
	}
	result := [][]int{}
	for i := 0; i <= total; i++ {
		subResList := allArranges(total-i, parts-1)
		for j := range subResList {
			result = append(result, append([]int{i}, subResList[j]...))
		}
	}

	return result
}

func currentIsValid(c Constraint, current []bool) bool {
	group := []int{}
	currentCount := 0
	for _, b := range current {
		if b {
			currentCount++
		} else {
			if currentCount > 0 {
				group = append(group, currentCount)
			}
			currentCount = 0
		}
	}
	if currentCount > 0 {
		group = append(group, currentCount)
	}

	if len(c) < len(group) {
		return false
	}

	for i := range group {
		if c[i] < group[i] {
			return false
		}
	}

	return true
}

func (p *Puzzle) getColumn(index int) []bool {
	bs := []bool{}
	for i := range p.Grids {
		bs = append(bs, p.Grids[i][index])
	}
	return bs
}

// after try fill a new row, validate it by checking all columns
func (p *Puzzle) IsValidForEachColumns() bool {
	for i := range p.TopConstraints {
		if !currentIsValid(p.TopConstraints[i], p.getColumn(i)) {
			return false
		}
	}
	return true
}

func (p *Puzzle) recursiveResolve(rowIndex int) {
	if rowIndex >= len(p.LeftConstraints) {
		return
	}
	possibles := GetPossiblesRows(len(p.Grids[0]), p.LeftConstraints[rowIndex])
	for pi := range possibles {
		row := possibles[pi]
		raw := p.Grids[rowIndex]
		p.Grids[rowIndex] = row
		if p.IsValidForEachColumns() {
			p.recursiveResolve(rowIndex + 1)
		} else {
			p.Grids[rowIndex] = raw
		}
	}
}

func (p *Puzzle) Resolve() {
	p.recursiveResolve(0)
	if p.IsValidForEachColumns() {
		fmt.Println("All Done")
	} else {
		fmt.Println("Oops, There was a bug.")
	}

}

func NewPuzzle(columns, rows int, left, top Constraints) Puzzle {
	grids := [][]bool{}
	for r := 0; r < rows; r++ {
		r := []bool{}
		for c := 0; c < columns; c++ {
			r = append(r, false)
		}
		grids = append(grids, r)
	}
	return Puzzle{
		Grids:           grids,
		LeftConstraints: left,
		TopConstraints:  top,
	}
}
