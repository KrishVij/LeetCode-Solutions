package main

import (
	"fmt"
	"sort"
)

func chopFromLast(arr []int) []int {

	if len(arr) <= 0 {

		return []int{}
	}

	return arr[:len(arr)-1]
}

func findAlluniqueSubsets(candidates []int, subsetStor []int, result *[][]int, target int, index int) {

	if target == 0 {

		subsetCopy := append([]int{}, subsetStor...)
		*result = append(*result, subsetCopy)

		return
	}

	for i := index; i < len(candidates); i++ {

		if i > index && candidates[i] == candidates[i-1] {

			continue
		}
		if candidates[i] > target {

			break
		}

		subsetStor = append(subsetStor, candidates[i])
		findAlluniqueSubsets(candidates, subsetStor, result, target-candidates[i], i+1)
		subsetStor = chopFromLast(subsetStor)
	}
}

func combinationSum2(candidates []int, target int) [][]int {

	var result [][]int
	sort.Ints(candidates)

	findAlluniqueSubsets(candidates, []int{}, &result, target, 0)

	return result
}
