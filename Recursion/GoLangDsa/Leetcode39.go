package main

import "fmt"

func chopFromLast(arr []int) []int {

	if len(arr) <= 0 {

		return []int{0}
	}

	return arr[:len(arr)-1]
}

func calculateAllCandidates(candidates []int, candidateStore []int, result *[][]int, target int, index int) {

	if target == 0 {

		subSetcopy := append([]int{}, candidateStore...)
		*result = append(*result, subSetcopy)
		return

	}
	if index >= len(candidates) || target < 0 {

		return
	}
	candidateStore = append(candidateStore, candidates[index])
	if candidates[index] <= target {

		calculateAllCandidates(candidates, candidateStore, result, target-candidates[index], index)
	} else {

		candidateStore = chopFromLast(candidateStore)
		calculateAllCandidates(candidates, candidateStore, result, target, index+1)
	}

}

func combinationSum(candidates []int, target int) [][]int {

	var result [][]int

	calculateAllCandidates(candidates, []int{}, &result, target, 0)
	return result
}
