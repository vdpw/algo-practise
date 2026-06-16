package archive

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/3sum/description/
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		// Smallest possible sum for this i is already > 0.
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}

		// Largest possible sum for this i is still < 0.
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return res
}

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	split := func(i int, nums []int) (int, int) {
		l, r := math.MinInt, math.MaxInt
		if i > 0 {
			l = nums[i-1]
		}
		if i < len(nums) {
			r = nums[i]
		}
		return l, r
	}

	max := func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}

	min := func(v1, v2 int) int {
		if v1 < v2 {
			return v1
		}
		return v2
	}

	find := func(l1, l2, l1from, l1end int) (int, int) {
		i := (l1from + l1end) / 2
		return i, (l1+l2+1)/2 - i

	}

	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		nums1, nums2 = nums2, nums1
		l1, l2 = l2, l1
	}
	l1from, l1end := 0, l1
	for l1from <= l1end {
		i, j := find(l1, l2, l1from, l1end)
		left1, right1 := split(i, nums1)
		left2, right2 := split(j, nums2)
		if left1 > right2 {
			l1end = i - 1
		} else if left2 > right1 {
			l1from = i + 1
		} else {
			if (l1+l2)%2 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2.0
			}
			return float64(max(left1, left2))
		}
	}
	return 0
}
