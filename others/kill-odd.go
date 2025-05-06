package others

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(max int) []int {
	arr := []int{}

	for i := 1; i <= max; i++ {
		arr = append(arr, i)
	}

	return arr
}

func doKillAllOdds(arr []int) []int {
	remains := []int{}
	for i, n := range arr {
		if i&1 != 0 {
			remains = append(remains, n)
		}
	}
	return remains
}

func RunKillOdd(mount int) {
	numbers := generateNumbers(mount)
	round := 1
	for {
		numbers = doKillAllOdds(numbers)
		fmt.Printf("round: %d, remains: %v\n", round, numbers)
		round++
		if len(numbers) <= 1 {
			break
		}
	}
}

func doKillRandomOdd(numbers []int, r *rand.Rand) []int {
	randIdx := r.Intn(len(numbers))
	if randIdx%2 == 1 {
		randIdx = randIdx - 1
	}
	result := []int{}
	result = append(result, numbers[:randIdx]...)
	result = append(result, numbers[randIdx+1:]...)
	return result
}

func RunKillRandom(mount int) {
	t := 100
	for {
		if t <= 0 {
			break
		}
		r := rand.New(rand.NewSource(time.Now().UnixMicro()))
		numbers := generateNumbers(mount)
		round := 1
		for {
			if len(numbers) <= 1 {
				break
			}
			round++
			numbers = doKillRandomOdd(numbers, r)

		}
		fmt.Printf("round: %d, alive: %v\n", round, numbers)
		t--
	}

}
