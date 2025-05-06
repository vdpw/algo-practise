package archive

var numbersToPut = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solveSudoku(board [][]byte) {
	b := NewBoard(board)
	b.Solve()
}

type Board struct {
	Panel      [][]byte
	Rows       [9]int
	Cols       [9]int
	Boxes      [9]int
	EmptyCells [][2]int
}

func NewBoard(p [][]byte) *Board {
	b := &Board{Panel: p}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if p[i][j] != '.' {
				num := p[i][j] - '1'
				mask := 1 << num
				b.Rows[i] |= mask
				b.Cols[j] |= mask
				boxIndex := (i/3)*3 + j/3
				b.Boxes[boxIndex] |= mask
			} else {
				b.EmptyCells = append(b.EmptyCells, [2]int{i, j})
			}
		}
	}
	return b
}

func (b *Board) Solve() bool {
	return b.solveHelper(0)
}

func (b *Board) solveHelper(index int) bool {
	if index == len(b.EmptyCells) {
		return true
	}
	row, column := b.EmptyCells[index][0], b.EmptyCells[index][1]
	boxIndex := (row/3)*3 + column/3

	for _, v := range numbersToPut {
		num := v - '1'
		mask := 1 << num
		if (b.Rows[row]&mask) == 0 && (b.Cols[column]&mask) == 0 && (b.Boxes[boxIndex]&mask) == 0 {
			b.Panel[row][column] = v
			b.Rows[row] |= mask
			b.Cols[column] |= mask
			b.Boxes[boxIndex] |= mask

			if b.solveHelper(index + 1) {
				return true
			}

			// backtracking
			b.Panel[row][column] = '.'
			b.Rows[row] &^= mask
			b.Cols[column] &^= mask
			b.Boxes[boxIndex] &^= mask
		}
	}
	return false
}

// https://leetcode.cn/problems/valid-sudoku/  有效但不需要一定有解
func isValidSudoku(board [][]byte) bool {
	Rows := [9]int{}
	Cols := [9]int{}
	Boxes := [9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				mask := 1 << num
				boxIndex := (i/3)*3 + j/3

				if (Rows[i]&mask) == 0 && (Cols[j]&mask) == 0 && (Boxes[boxIndex]&mask) == 0 {
					Rows[i] |= mask
					Cols[j] |= mask
					Boxes[boxIndex] |= mask
				} else {
					return false
				}
			}
		}
	}
	return true

}
