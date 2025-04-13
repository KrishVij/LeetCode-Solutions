package main

func chopFromLast(nums []int) []int {

	if len(nums) <= 0 {

		return []int{0}
	}

	return nums[:len(nums)-1]
}

func calculateAllSubsets(nums, subsetStore []int, index int, result *[][]int) {

	if index >= len(nums) {

		subsetCopy := append([]int{}, subsetStore...)
		*result = append(*result, subsetCopy)
		return
	}

	subsetStore = append(subsetStore, nums[index])
	calculateAllSubsets(nums, subsetStore, index+1, result)

	subsetStore = chopFromLast(subsetStore)
	calculateAllSubsets(nums, subsetStore, index+1, result)

}

func subsets(nums []int) [][]int {

	var result [][]int
	calculateAllSubsets(nums, []int{}, 0, &result)
	return result
}
