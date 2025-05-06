package others

import (
	"math"
	"strconv"
)

func analyze(conditionNumber []int, digitWithPos map[int]int) (int, int) {
	numMatchCount := 0
	posMatchCount := 0
	for i := range conditionNumber {
		pos, found := digitWithPos[conditionNumber[i]]
		if found {
			numMatchCount++
			if pos == i {
				posMatchCount++
			}
		}
	}

	return numMatchCount, posMatchCount
}

func intToArr(n int) []int {
	result := []int{}
	s := strconv.Itoa(n)
	for i := 0; i < len(s); i++ {
		result = append(result, int(s[i]-'0'))
	}

	return result
}

type Condition interface {
	checkNumbers([]int) bool
	checkDigitPositions(digitWithPos map[int]int) bool
}

type UniqNumCondition struct {
}

func (*UniqNumCondition) checkDigitPositions(map[int]int) bool { return true }
func (u *UniqNumCondition) checkNumbers(number []int) bool {
	set := map[int]struct{}{}
	for i := range number {
		set[number[i]] = struct{}{}
	}
	return len(set) == len(number)
}

type NumAndPosCondition struct {
	MatchNum int
	MatchPos int

	Data []int
}

func (*NumAndPosCondition) checkNumbers([]int) bool { return true }
func (c *NumAndPosCondition) checkDigitPositions(digitWithPos map[int]int) bool {
	n, p := analyze(c.Data, digitWithPos)
	return n == c.MatchNum && p == c.MatchPos
}

func GuessNumbers() []int {
	conditions := []Condition{
		&UniqNumCondition{},
		&NumAndPosCondition{1, 0, intToArr(9285)},
		&NumAndPosCondition{2, 0, intToArr(1937)},
		&NumAndPosCondition{1, 1, intToArr(5201)},
		&NumAndPosCondition{0, 0, intToArr(6507)},
	}

	result := []int{}
outer:
	for i := 1000; i <= 9999; i++ {
		arr := intToArr(i)
		digitWithPos := map[int]int{}
		for i := range arr {
			digitWithPos[arr[i]] = i
		}
		for j := range conditions {
			cond := conditions[j]
			if !cond.checkNumbers(arr) || !cond.checkDigitPositions(digitWithPos) {
				continue outer
			}
		}
		result = append(result, i)
	}

	return result
}

// x^2 + y^2 = 19451945
func GetTwoNumbers(target int) [][]int {
	max := int(math.Sqrt(float64(target))) + 1
	set := map[int]struct{}{}
	x := 1
	results := [][]int{}
	for ; x < max; x++ {
		if _, found := set[x]; found {
			continue
		}
		xp := x * x
		y := int(math.Sqrt(float64(target-xp))) - 1
		for ; y < max; y++ {
			if _, found := set[y]; found {
				continue
			}
			yp := y * y
			if yp+xp > target {
				break
			}
			if yp+xp == target {
				set[x] = struct{}{}
				set[y] = struct{}{}
				results = append(results, []int{x, y})
			}
		}
	}
	return results
}
