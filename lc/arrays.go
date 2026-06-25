package lc

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
