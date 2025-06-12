package nonograms

import (
	"fmt"
)

// -----------------------------------------------------------------------------
// BASIC TYPES
// -----------------------------------------------------------------------------

const N = 50             // board edge (change to 10,15… while testing)
type Pattern = uint64    // one bitboard: bit i == cell i filled
type LineSet = []Pattern // all legal patterns for a line
type GridRow = Pattern   // current committed cells of a row
type Grid = [N]GridRow   // committed rows of the whole board

// -----------------------------------------------------------------------------
// 1.  PATTERN GENERATION  ------------------------------------------------------
// -----------------------------------------------------------------------------

// generatePatterns returns every bit pattern of length 'size' that satisfies
// the clue, e.g. clue [3 2] on a 7-cell line ->  111.11..
func generatePatterns(size int, clue []int) LineSet {
	var out LineSet
	var place func(pos, clueIdx int, p Pattern)

	place = func(pos, clueIdx int, p Pattern) {
		if clueIdx == len(clue) { // no more blocks -> fill rest with 0
			if pos <= size { // valid line
				out = append(out, p)
			}
			return
		}
		blk := clue[clueIdx]
		// leftmost position where this block may start
		for start := pos; start+blk <= size; start++ {
			// add block
			var q = p
			for i := 0; i < blk; i++ {
				q |= 1 << (size - 1 - (start + i))
			}
			// one empty cell after the block (unless it ends the line)
			nextPos := start + blk + 1
			if nextPos > size { // last block fits exactly
				if clueIdx == len(clue)-1 {
					out = append(out, q)
				}
				continue
			}
			place(nextPos, clueIdx+1, q)
		}
	}
	place(0, 0, 0)
	return out
}

// -----------------------------------------------------------------------------
// 2.  CONSTRAINT PROPAGATION LOOP  --------------------------------------------
// -----------------------------------------------------------------------------

// forcedBits returns two bitmasks:
//
//	filled = bits that are 1 in *every* pattern
//	empty  = bits that are 0 in *every* pattern
func forcedBits(set LineSet) (filled, empty Pattern) {
	if len(set) == 0 {
		return 0, 0
	}
	filled = ^Pattern(0) // all 1s
	empty = ^filled      // all 0s
	for _, p := range set {
		filled &= p
		empty &= ^p
	}
	return
}

// propagate removes impossible patterns until nothing changes.
// It returns false if a contradiction is found.
func propagate(rowSet, colSet []LineSet, grid *Grid) bool {
	type item struct {
		isRow bool
		idx   int
	}
	queue := make([]item, 0, 2*N)
	for i := 0; i < N; i++ { // everything dirty once
		queue = append(queue, item{true, i}, item{false, i})
	}

	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]

		if it.isRow {
			f, e := forcedBits(rowSet[it.idx])
			r := grid[it.idx]
			if (r&f) != f || (r&^e) != r { // row gained info
				grid[it.idx] = (r | f) & ^e
				// every affected column becomes dirty
				for c := 0; c < N; c++ {
					queue = append(queue, item{false, c})
				}
			}
			// keep only row patterns that agree with committed bits
			rowSet[it.idx] = filter(rowSet[it.idx], grid[it.idx])
			if len(rowSet[it.idx]) == 0 {
				return false
			}
		} else { // column
			colBits := columnBits(*grid, it.idx)
			colSet[it.idx] = filter(colSet[it.idx], colBits)
			if len(colSet[it.idx]) == 0 {
				return false
			}
			f, e := forcedBits(colSet[it.idx])
			for r := 0; r < N; r++ {
				bit := Pattern(1) << (N - 1 - r)
				changed := false
				if f&bit != 0 && grid[r]&bit == 0 {
					grid[r] |= bit
					changed = true
				}
				if e&bit != 0 && grid[r]&bit != 0 {
					grid[r] &^= bit
					changed = true
				}
				if changed {
					queue = append(queue, item{true, r})
				}
			}
		}
	}
	return true
}

// columnBits extracts the rth *committed* column bits into a pattern
func columnBits(g Grid, col int) Pattern {
	var p Pattern
	for r := 0; r < N; r++ {
		if g[r]&(1<<(N-1-col)) != 0 {
			p |= 1 << (N - 1 - r)
		}
	}
	return p
}

// filter keeps only patterns compatible with committed bits (=known 1s/0s)
func filter(set LineSet, committed Pattern) LineSet {
	out := set[:0]
	mask := committed != 0
	for _, p := range set {
		if mask && ((p & committed) != committed) {
			continue // violates a known 1
		}
		if mask && ((^p)&committed) != 0 {
			continue // violates a known 0
		}
		out = append(out, p)
	}
	return out
}

// -----------------------------------------------------------------------------
// 3.  DFS WITH MINIMUM-BRANCH HEURISTIC  --------------------------------------
// -----------------------------------------------------------------------------

func solve(rowSet, colSet []LineSet, grid *Grid) (Grid, bool) {
	if !propagate(rowSet, colSet, grid) {
		return *grid, false
	}
	// check if solved
	done := true
	for _, s := range rowSet {
		if len(s) != 1 {
			done = false
			break
		}
	}
	if done {
		return *grid, true
	}
	// choose the line with fewest patterns > 1
	minIdx, isRow := -1, true
	minSize := 1 << 30
	for i, s := range rowSet {
		if l := len(s); l > 1 && l < minSize {
			minIdx, minSize, isRow = i, l, true
		}
	}
	for i, s := range colSet {
		if l := len(s); l > 1 && l < minSize {
			minIdx, minSize, isRow = i, l, false
		}
	}

	// branch over its patterns
	if isRow {
		old := append(LineSet(nil), rowSet[minIdx]...)
		for _, p := range old {
			rowSet[minIdx] = LineSet{p}
			clone := *grid
			if g, ok := solve(rowSet, colSet, &clone); ok {
				return g, true
			}
		}
		rowSet[minIdx] = old
	} else {
		old := append(LineSet(nil), colSet[minIdx]...)
		for _, p := range old {
			colSet[minIdx] = LineSet{p}
			clone := *grid
			if g, ok := solve(rowSet, colSet, &clone); ok {
				return g, true
			}
		}
		colSet[minIdx] = old
	}
	return *grid, false
}

// -----------------------------------------------------------------------------
// 4.  DEMO DRIVER  -------------------------------------------------------------
// -----------------------------------------------------------------------------

func main() {
	// Toy 5×5 sample first (easy to see)
	left := [][]int{
		{1, 1}, {5}, {1, 1}, {5}, {1, 1},
	}
	top := [][]int{
		{1, 1}, {5}, {1, 1}, {5}, {1, 1},
	}
	out, ok := Run(left, top)
	fmt.Println("Solved small:", ok)
	printGrid(out)

	// Real 50×50: replace with your own clues
	// left50 := make([][]int, N) // fill with real data
	// top50 := make([][]int, N)
	// out, ok = Run(left50, top50)
	// fmt.Println("Solved big:", ok)
}

// Run wires everything together
func Run(left, top [][]int) (Grid, bool) {
	var grid Grid
	rowSet := make([]LineSet, N)
	colSet := make([]LineSet, N)
	for i := 0; i < N; i++ {
		rowSet[i] = generatePatterns(N, left[i])
		colSet[i] = generatePatterns(N, top[i])
	}
	return solve(rowSet, colSet, &grid)
}

// printGrid prints ▓/·
func printGrid(g Grid) {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if g[r]&(1<<(N-1-c)) != 0 {
				fmt.Print("▓")
			} else {
				fmt.Print("·")
			}
		}
		fmt.Println()
	}
}
