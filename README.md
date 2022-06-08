# golang_combination_sum_v2

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sum to `target`.

Each number in `candidates` may only be used **once** in the combination.

**Note:** The solution set must not contain duplicate combinations.

## Examples

**Example 1:**

```
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

```

**Example 2:**

```
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

```

**Constraints:**

- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

## 解析

題目給了一個可能會出現重複值的整數陣列 nums, 還有一個整數 target

要求寫出一個演算法找出所有和 = target 的子陣列不重複組合 

這題與 [39. Combination Sum](https://www.notion.so/39-Combination-Sum-f2f76b649cce4a66b480ad526b933bbd) 類似

差別在於

1. 一個元素只能使用一次
2. 組合不能重複

元素使用一次只要使用類似 PowerSet 的作法即可

把每個元素分為選擇或不選擇兩種可能

組合要不能重複

就必須要把重複元素分為 選擇1次以上或者不選

要照這種方法來做到必須先把陣列做排序

然後每次只要發現前一個元素與當下元素相同就 skip到下一個

舉例來說: 給定 candidates = [10,1,2,7,6,1,5], target = 8

先把 candidates 排序

變成 [1,1,2,5,6,7,10]

![](https://i.imgur.com/qVRh0gi.png)

透過排序過的陣列逐步檢查出所有可能的組合

每個數字都有選擇與不選擇兩種可能所是 $O(2^n)$

## 程式碼
```go
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

```
## 困難點

1. 理解如何窮舉
2. 理解如何避免重複

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity