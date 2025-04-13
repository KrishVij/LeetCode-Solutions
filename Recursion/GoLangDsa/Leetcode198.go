package main

import "fmt"

func stealingFromHouse(nums []int, index int, money int) int {

	if index >= len(nums) {

		return money
	}

	return stealingFromHouse(nums, index+2, money+nums[index])
}

func rob(nums []int) int {

	return stealingFromHouse(nums, 0, 0)
}
