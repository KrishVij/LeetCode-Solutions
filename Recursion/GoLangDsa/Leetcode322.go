package main

import "fmt"

func findFewestCoinsForAmount(nums []int, target, count, index int) int {

	if target == 0 {

		return count + 1
	}

	if index >= len(nums) || target < 0 {

		return -1
	}

	if nums[index] <= target {

		findFewestCoinsForAmount(nums, target-nums[index], count, index)
	} else {

		findFewestCoinsForAmount(nums, target, count, index+1)
	}

	return -1

}

func coinChange(coins []int, amount int) int {

	return findFewestCoinsForAmount(coins, amount, 0, 0)
}
