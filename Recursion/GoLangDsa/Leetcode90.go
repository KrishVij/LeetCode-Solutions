package main

import "sort"

func chopFromLast(nums []int) []int {

	if len(nums) <= 0 {

		return []int{0}
	}

	return nums[:len(nums)-1]
}

func calculateAllSubsets(nums, subsetStore []int, index int, result *[][]int) [][]int {

	*result = append(*result, append([]int{}, subsetStore...))

	for i := index; i < len(nums); i++ {

		if i != index && nums[i] == nums[i-1] {

			continue
		}

		subsetStore = append(subsetStore, nums[i])
		calculateAllSubsets(nums, subsetStore, i+1, result)
		subsetStore = chopFromLast(subsetStore)
	}

	return *result
}

func subsetsWithDup(nums []int) [][]int {

	var result [][]int
	var subsetStore []int

	sort.Ints(nums)

	calculateAllSubsets(nums, subsetStore, 0, &result)

	return result

}
