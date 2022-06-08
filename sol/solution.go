package sol

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// sort candidates
	sort.Ints(candidates)
	result := [][]int{}
	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) {
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			result = append(result, temp)
			return
		}
		if i >= len(candidates) || total > target {
			return
		}
		// choose i
		cur = append(cur, candidates[i])
		dfs(i+1, cur, total+candidates[i])
		cur = cur[:len(cur)-1]

		// not choose i
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(i+1, cur, total)
	}
	dfs(0, []int{}, 0)
	return result
}
