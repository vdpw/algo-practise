package main

import (
	"errors"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

const maxLineLength = 64

var (
	ErrNoSolution  = errors.New("nonograms: no solution")
	ErrSearchLimit = errors.New("nonograms: search node limit reached")
)

type Options struct {
	// MaxSearchNodes caps DFS nodes. Zero means no limit.
	MaxSearchNodes int
}

type Solution struct {
	Width    int
	Height   int
	Rows     []uint64
	RowClues [][]int
	ColClues [][]int
}

func (s Solution) At(row, col int) bool {
	return s.Rows[row]&(uint64(1)<<uint(col)) != 0
}

func (s Solution) BoolGrid() [][]bool {
	grid := make([][]bool, s.Height)
	for r := 0; r < s.Height; r++ {
		grid[r] = make([]bool, s.Width)
		for c := 0; c < s.Width; c++ {
			grid[r][c] = s.At(r, c)
		}
	}
	return grid
}

func (s Solution) String() string {
	if len(s.RowClues) != s.Height || len(s.ColClues) != s.Width {
		return s.GridString()
	}

	rowLabels := make([]string, s.Height)
	rowLabelWidth := 0
	for r, clue := range s.RowClues {
		rowLabels[r] = clueLabel(clue)
		if len(rowLabels[r]) > rowLabelWidth {
			rowLabelWidth = len(rowLabels[r])
		}
	}

	colLabels := make([][]string, s.Width)
	colDepth := 0
	cellWidth := 1
	for c, clue := range s.ColClues {
		colLabels[c] = clueParts(clue)
		if len(colLabels[c]) > colDepth {
			colDepth = len(colLabels[c])
		}
		for _, part := range colLabels[c] {
			if len(part) > cellWidth {
				cellWidth = len(part)
			}
		}
	}

	var b strings.Builder
	for line := 0; line < colDepth; line++ {
		if rowLabelWidth > 0 {
			b.WriteString(strings.Repeat(" ", rowLabelWidth))
			b.WriteByte(' ')
		}
		for c := 0; c < s.Width; c++ {
			if c > 0 {
				b.WriteByte(' ')
			}
			offset := colDepth - len(colLabels[c])
			if line < offset {
				b.WriteString(strings.Repeat(" ", cellWidth))
				continue
			}
			b.WriteString(leftPad(colLabels[c][line-offset], cellWidth))
		}
		b.WriteByte('\n')
	}

	for r := 0; r < s.Height; r++ {
		if r > 0 {
			b.WriteByte('\n')
		}
		if rowLabelWidth > 0 {
			b.WriteString(leftPad(rowLabels[r], rowLabelWidth))
			b.WriteByte(' ')
		}
		for c := 0; c < s.Width; c++ {
			if c > 0 {
				b.WriteByte(' ')
			}
			if s.At(r, c) {
				b.WriteString(leftPad("#", cellWidth))
			} else {
				b.WriteString(leftPad(".", cellWidth))
			}
		}
	}
	return b.String()
}

func (s Solution) GridString() string {
	var b strings.Builder
	for r := 0; r < s.Height; r++ {
		if r > 0 {
			b.WriteByte('\n')
		}
		for c := 0; c < s.Width; c++ {
			if s.At(r, c) {
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
		}
	}
	return b.String()
}

func clueLabel(clue []int) string {
	return strings.Join(clueParts(clue), " ")
}

func clueParts(clue []int) []string {
	if len(clue) == 0 {
		return []string{"0"}
	}

	parts := make([]string, len(clue))
	for i, n := range clue {
		parts[i] = strconv.Itoa(n)
	}
	return parts
}

func leftPad(text string, width int) string {
	if len(text) >= width {
		return text
	}
	return strings.Repeat(" ", width-len(text)) + text
}

func Solve(rowClues, colClues [][]int) (Solution, error) {
	return SolveWithOptions(rowClues, colClues, Options{})
}

func SolveWithOptions(rowClues, colClues [][]int, opts Options) (Solution, error) {
	st, err := newState(rowClues, colClues)
	if err != nil {
		return Solution{}, err
	}

	s := solver{opts: opts}
	solution, ok, err := s.search(st)
	if err != nil {
		return Solution{}, err
	}
	if !ok {
		return Solution{}, ErrNoSolution
	}
	return solution, nil
}

type solver struct {
	opts  Options
	nodes int
}

func (s *solver) search(st *state) (Solution, bool, error) {
	s.nodes++
	if s.opts.MaxSearchNodes > 0 && s.nodes > s.opts.MaxSearchNodes {
		return Solution{}, false, ErrSearchLimit
	}

	if err := st.propagate(); err != nil {
		if errors.Is(err, ErrNoSolution) {
			return Solution{}, false, nil
		}
		return Solution{}, false, err
	}
	if st.complete() {
		return st.solution(), true, nil
	}

	branch, ok := st.chooseBranch()
	if !ok {
		return st.solution(), true, nil
	}

	var candidates []uint64
	if branch.isRow {
		candidates = append(candidates, st.rowSets[branch.idx]...)
	} else {
		candidates = append(candidates, st.colSets[branch.idx]...)
	}

	for _, pattern := range candidates {
		next := st.clone()
		if branch.isRow {
			next.rowSets[branch.idx] = []uint64{pattern}
		} else {
			next.colSets[branch.idx] = []uint64{pattern}
		}

		solution, solved, err := s.search(next)
		if err != nil {
			return Solution{}, false, err
		}
		if solved {
			return solution, true, nil
		}
	}

	return Solution{}, false, nil
}

type state struct {
	width  int
	height int

	rowMask uint64
	colMask uint64

	rowSets  [][]uint64
	colSets  [][]uint64
	rowClues [][]int
	colClues [][]int

	filledRows []uint64
	emptyRows  []uint64
}

type lineRef struct {
	isRow bool
	idx   int
}

func newState(rowClues, colClues [][]int) (*state, error) {
	height := len(rowClues)
	width := len(colClues)
	if height == 0 || width == 0 {
		return nil, fmt.Errorf("nonograms: board must have at least one row and one column")
	}
	if height > maxLineLength || width > maxLineLength {
		return nil, fmt.Errorf("nonograms: max supported line length is %d", maxLineLength)
	}

	rows, rowTotal, err := normalizeClues(rowClues, width, "row")
	if err != nil {
		return nil, err
	}
	cols, colTotal, err := normalizeClues(colClues, height, "column")
	if err != nil {
		return nil, err
	}
	if rowTotal != colTotal {
		return nil, fmt.Errorf("nonograms: row clues fill %d cells but column clues fill %d cells", rowTotal, colTotal)
	}

	rowSets := make([][]uint64, height)
	for r, clue := range rows {
		rowSets[r] = generatePatterns(width, clue)
	}

	colSets := make([][]uint64, width)
	for c, clue := range cols {
		colSets[c] = generatePatterns(height, clue)
	}

	return &state{
		width:      width,
		height:     height,
		rowMask:    lineMask(width),
		colMask:    lineMask(height),
		rowSets:    rowSets,
		colSets:    colSets,
		rowClues:   cloneClues(rows),
		colClues:   cloneClues(cols),
		filledRows: make([]uint64, height),
		emptyRows:  make([]uint64, height),
	}, nil
}

func (s *state) propagate() error {
	queue := make([]lineRef, 0, s.height+s.width)
	rowQueued := make([]bool, s.height)
	colQueued := make([]bool, s.width)

	pushRow := func(row int) {
		if !rowQueued[row] {
			rowQueued[row] = true
			queue = append(queue, lineRef{isRow: true, idx: row})
		}
	}
	pushCol := func(col int) {
		if !colQueued[col] {
			colQueued[col] = true
			queue = append(queue, lineRef{idx: col})
		}
	}

	for r := 0; r < s.height; r++ {
		pushRow(r)
	}
	for c := 0; c < s.width; c++ {
		pushCol(c)
	}

	for head := 0; head < len(queue); head++ {
		item := queue[head]
		if item.isRow {
			rowQueued[item.idx] = false
			if err := s.propagateRow(item.idx, pushCol); err != nil {
				return err
			}
		} else {
			colQueued[item.idx] = false
			if err := s.propagateCol(item.idx, pushRow); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *state) propagateRow(row int, pushCol func(int)) error {
	filtered := filterPatterns(s.rowSets[row], s.filledRows[row], s.emptyRows[row])
	if len(filtered) == 0 {
		return ErrNoSolution
	}
	s.rowSets[row] = filtered

	forcedFilled, forcedEmpty := forcedBits(filtered, s.rowMask)
	if forcedFilled&s.emptyRows[row] != 0 || forcedEmpty&s.filledRows[row] != 0 {
		return ErrNoSolution
	}

	oldFilled := s.filledRows[row]
	oldEmpty := s.emptyRows[row]
	s.filledRows[row] |= forcedFilled
	s.emptyRows[row] |= forcedEmpty

	changed := (oldFilled ^ s.filledRows[row]) | (oldEmpty ^ s.emptyRows[row])
	for changed != 0 {
		col := bits.TrailingZeros64(changed)
		pushCol(col)
		changed &^= uint64(1) << uint(col)
	}
	return nil
}

func (s *state) propagateCol(col int, pushRow func(int)) error {
	filled, empty := s.columnKnown(col)
	filtered := filterPatterns(s.colSets[col], filled, empty)
	if len(filtered) == 0 {
		return ErrNoSolution
	}
	s.colSets[col] = filtered

	forcedFilled, forcedEmpty := forcedBits(filtered, s.colMask)
	forced := forcedFilled | forcedEmpty
	for forced != 0 {
		row := bits.TrailingZeros64(forced)
		rowBit := uint64(1) << uint(row)
		cellFilled := forcedFilled&rowBit != 0
		changed, err := s.setCell(row, col, cellFilled)
		if err != nil {
			return err
		}
		if changed {
			pushRow(row)
		}
		forced &^= rowBit
	}

	return nil
}

func (s *state) columnKnown(col int) (filled, empty uint64) {
	colBit := uint64(1) << uint(col)
	for r := 0; r < s.height; r++ {
		rowBit := uint64(1) << uint(r)
		if s.filledRows[r]&colBit != 0 {
			filled |= rowBit
		}
		if s.emptyRows[r]&colBit != 0 {
			empty |= rowBit
		}
	}
	return filled, empty
}

func (s *state) setCell(row, col int, filled bool) (bool, error) {
	bit := uint64(1) << uint(col)
	if filled {
		if s.emptyRows[row]&bit != 0 {
			return false, ErrNoSolution
		}
		if s.filledRows[row]&bit != 0 {
			return false, nil
		}
		s.filledRows[row] |= bit
		return true, nil
	}

	if s.filledRows[row]&bit != 0 {
		return false, ErrNoSolution
	}
	if s.emptyRows[row]&bit != 0 {
		return false, nil
	}
	s.emptyRows[row] |= bit
	return true, nil
}

func (s *state) complete() bool {
	for r := 0; r < s.height; r++ {
		if (s.filledRows[r]|s.emptyRows[r])&s.rowMask != s.rowMask {
			return false
		}
	}
	return true
}

func (s *state) chooseBranch() (lineRef, bool) {
	best := lineRef{}
	bestSize := int(^uint(0) >> 1)

	for r, candidates := range s.rowSets {
		if size := len(candidates); size > 1 && size < bestSize {
			best = lineRef{isRow: true, idx: r}
			bestSize = size
		}
	}
	for c, candidates := range s.colSets {
		if size := len(candidates); size > 1 && size < bestSize {
			best = lineRef{idx: c}
			bestSize = size
		}
	}

	return best, bestSize != int(^uint(0)>>1)
}

func (s *state) clone() *state {
	return &state{
		width:      s.width,
		height:     s.height,
		rowMask:    s.rowMask,
		colMask:    s.colMask,
		rowSets:    clonePatternSets(s.rowSets),
		colSets:    clonePatternSets(s.colSets),
		rowClues:   cloneClues(s.rowClues),
		colClues:   cloneClues(s.colClues),
		filledRows: append([]uint64(nil), s.filledRows...),
		emptyRows:  append([]uint64(nil), s.emptyRows...),
	}
}

func (s *state) solution() Solution {
	return Solution{
		Width:    s.width,
		Height:   s.height,
		Rows:     append([]uint64(nil), s.filledRows...),
		RowClues: cloneClues(s.rowClues),
		ColClues: cloneClues(s.colClues),
	}
}

func clonePatternSets(src [][]uint64) [][]uint64 {
	dst := make([][]uint64, len(src))
	for i := range src {
		dst[i] = append([]uint64(nil), src[i]...)
	}
	return dst
}

func cloneClues(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = append([]int(nil), src[i]...)
	}
	return dst
}

func normalizeClues(clues [][]int, lineLength int, axis string) ([][]int, int, error) {
	normalized := make([][]int, len(clues))
	totalFilled := 0

	for i, clue := range clues {
		blocks := make([]int, 0, len(clue))
		lineTotal := 0
		for _, n := range clue {
			if n < 0 {
				return nil, 0, fmt.Errorf("nonograms: %s %d has negative clue %d", axis, i, n)
			}
			if n == 0 {
				continue
			}
			blocks = append(blocks, n)
			lineTotal += n
		}

		minLength := lineTotal
		if len(blocks) > 1 {
			minLength += len(blocks) - 1
		}
		if minLength > lineLength {
			return nil, 0, fmt.Errorf("nonograms: %s %d clue needs %d cells in a %d-cell line", axis, i, minLength, lineLength)
		}

		normalized[i] = blocks
		totalFilled += lineTotal
	}

	return normalized, totalFilled, nil
}

func generatePatterns(length int, clue []int) []uint64 {
	if len(clue) == 0 {
		return []uint64{0}
	}

	minSuffix := make([]int, len(clue)+1)
	for i := len(clue) - 1; i >= 0; i-- {
		minSuffix[i] = minSuffix[i+1] + clue[i]
		if i < len(clue)-1 {
			minSuffix[i]++
		}
	}

	patterns := make([]uint64, 0)
	var place func(clueIdx, pos int, pattern uint64)
	place = func(clueIdx, pos int, pattern uint64) {
		if clueIdx == len(clue) {
			patterns = append(patterns, pattern)
			return
		}

		block := clue[clueIdx]
		latestStart := length - minSuffix[clueIdx]
		for start := pos; start <= latestStart; start++ {
			nextPattern := pattern | rangeMask(start, block)
			nextPos := start + block
			if clueIdx < len(clue)-1 {
				nextPos++
			}
			place(clueIdx+1, nextPos, nextPattern)
		}
	}

	place(0, 0, 0)
	return patterns
}

func filterPatterns(patterns []uint64, filled, empty uint64) []uint64 {
	out := patterns[:0]
	for _, pattern := range patterns {
		if pattern&filled == filled && pattern&empty == 0 {
			out = append(out, pattern)
		}
	}
	return out
}

func forcedBits(patterns []uint64, mask uint64) (filled, empty uint64) {
	filled = mask
	var possibleFilled uint64
	for _, pattern := range patterns {
		filled &= pattern
		possibleFilled |= pattern
	}
	empty = (^possibleFilled) & mask
	return filled, empty
}

func lineMask(length int) uint64 {
	if length == maxLineLength {
		return ^uint64(0)
	}
	return (uint64(1) << uint(length)) - 1
}

func rangeMask(start, length int) uint64 {
	return lineMask(length) << uint(start)
}
