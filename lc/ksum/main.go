package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return kSum(nums, 0, 3, 0)
}

// NOTE: hard to learn
// kSum returns all k-tuples (ascending) that sum to target.
func kSum(nums []int, target, k, start int) [][]int {
	n := len(nums)
	res := [][]int{}

	// Base case: 2-Sum with two-pointers
	if k == 2 {
		lo, hi := start, n-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			switch {
			case sum < target:
				lo++
			case sum > target:
				hi--
			default:
				res = append(res, []int{nums[lo], nums[hi]})
				lo, hi = lo+1, hi-1
				for lo < hi && nums[lo] == nums[lo-1] { // dup skip
					lo++
				}
				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
			}
		}
		return res
	}

	// Recursive case: fix nums[i], solve (k-1)-Sum
	for i := start; i <= n-k; i++ {
		if i > start && nums[i] == nums[i-1] { // dup skip
			continue
		}

		// Early-exit pruning
		minSum := nums[i] + sumSlice(nums, i+1, i+k) // smallest k-1
		maxSum := nums[i] + sumSlice(nums, n-k+1, n) // largest k-1
		if target < minSum {
			break // too small no matter what
		}
		if target > maxSum {
			continue // need bigger nums[i]
		}

		sub := kSum(nums, target-nums[i], k-1, i+1)
		for _, t := range sub {
			res = append(res, append([]int{nums[i]}, t...))
		}
	}
	return res
}

// helper: sumSlice sums nums[l:r] (half-open)
func sumSlice(nums []int, l, r int) (s int) {
	for _, v := range nums[l:r] {
		s += v
	}
	return
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// Output: [[-1 -1 2] [-1 0 1]]
}
