package main

func calculateALLCombinations(nums []int, target int, sum int, index int) int {

	if index == len(nums) {

		if sum == target {

			return 1
		}

		return 0
	}

	count := calculateALLCombinations(nums, target, sum+nums[index], index+1)

	count += calculateALLCombinations(nums, target, sum-nums[index], index+1)

	return count
}

func findTargetSumWays(nums []int, target int) int {

	return calculateALLCombinations(nums, target, 0, 0)
}
