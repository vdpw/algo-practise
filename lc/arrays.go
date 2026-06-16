package lc

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
