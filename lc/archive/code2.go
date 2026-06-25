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

// https://leetcode.cn/problems/next-permutation/
func nextPermutation(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	p := l - 2
	for p >= 0 && nums[p+1] <= nums[p] {
		p--
	}
	if p >= 0 {
		successor := l - 1
		for nums[successor] <= nums[p] {
			successor--
		}
		nums[p], nums[successor] = nums[successor], nums[p]
	}
	for from, to := p+1, l-1; from < to; {
		nums[from], nums[to] = nums[to], nums[from]
		from++
		to--
	}
}

// https://leetcode.cn/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	if l < 3 {
		return target
	}
	distance := func(a, b int) int {
		d := a - b
		if d < 0 {
			return -d
		}
		return d
	}
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i <= l-2; i++ {
		left := i + 1
		right := l - 1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if distance(currentSum, target) < distance(closestSum, target) {
				closestSum = currentSum
			}
			v := currentSum - target
			if v > 0 {
				right--
			} else if v < 0 {
				left++
			} else {
				return target
			}
		}
	}
	return closestSum
}

// https://leetcode.cn/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	if l < 4 {
		return nil
	}
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < l-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := l - 1
			for left < right {
				currentSum := nums[i] + nums[j] + nums[left] + nums[right]
				if currentSum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if currentSum > target {
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else {
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
				}
			}
		}
	}
	return result
}
