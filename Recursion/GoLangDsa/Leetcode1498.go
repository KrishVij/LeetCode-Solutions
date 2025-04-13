package main

import "fmt"

const MOD = 1_000_000_007

func findAllSubsets(nums []int, target int, sum int, count int, index int) int {

	if index >= len(nums) {

		return count
	}

	sum += nums[index]
	if sum != target {

		return findAllSubsets(nums, target, sum, count, index+1) % MOD
	} else {

		return findAllSubsets(nums, target, sum, count+1, index+1) % MOD
	}
}

func numSubseq(nums []int, target int) int {

	return findAllSubsets(nums, target, 0, 0, 0)

}
