package lc

/*
4. 寻找两个正序数组的中位数
困难
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6
*/
// https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	indexMin, indexMax, halfLen := 0, m, (m+n+1)/2
	for indexMin <= indexMax {
		indexMiddle := (indexMin + indexMax) / 2
		j := halfLen - indexMiddle
		if indexMiddle < m && nums2[j-1] > nums1[indexMiddle] {
			indexMin = indexMiddle + 1
		} else if indexMiddle > 0 && nums1[indexMiddle-1] > nums2[j] {
			indexMax = indexMiddle - 1
		} else {
			var max_of_left, min_of_right int
			if indexMiddle == 0 {
				max_of_left = nums2[j-1]
			} else if j == 0 {
				max_of_left = nums1[indexMiddle-1]
			} else {
				max_of_left = nums1[indexMiddle-1]
				if nums2[j-1] > max_of_left {
					max_of_left = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(max_of_left)
			}
			if indexMiddle == m {
				min_of_right = nums2[j]
			} else if j == n {
				min_of_right = nums1[indexMiddle]
			} else {
				min_of_right = min(nums1[indexMiddle], nums2[j])
			}
			return float64(max_of_left+min_of_right) / 2.0
		}
	}
	return 0.0
}

/*
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。
*/
// https://leetcode.cn/problems/next-permutation/
func nextPermutation(nums []int) {

}

// https://leetcode.cn/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var dfs func(target int, path []int, start int)
	dfs = func(target int, path []int, start int) {
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if target < candidates[i] {
				continue
			}
			dfs(target-candidates[i], append(path, candidates[i]), i)
		}
	}
	dfs(target, []int{}, 0)
	return res
}

// https://leetcode.cn/problems/combination-sum-ii
func combinationSum2(candidates []int, target int) [][]int {
	counter := map[int]int{}
	for _, num := range candidates {
		counter[num]++
	}
	var uniqueCandidates [][]int
	for num, count := range counter {
		uniqueCandidates = append(uniqueCandidates, []int{num, count})
	}

	var res [][]int
	var dfs func(target int, path []int, start int, index int)
	dfs = func(target int, path []int, start int, index int) {
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		if start >= len(uniqueCandidates) {
			return
		}
		for i := start; i < len(uniqueCandidates); {
			num, count := uniqueCandidates[i][0], uniqueCandidates[i][1]
			if target < num {
				continue
			}
			if index >= count {
				i++
				continue
			}
			dfs(target-num, append(path, num), i, index+1)
		}
	}
	dfs(target, []int{}, 0, 0)
	return res
}
