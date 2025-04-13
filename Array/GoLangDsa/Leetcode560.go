package main

import "math"

func subarraySum(nums []int, k int) int {

	if k < len(nums) || k > len(nums) {

		return -1
	}
	count := 0

	for i := 0; i < len(nums); i++ {

		sum := 0
		for j := i; j < len(nums); j++ {

			sum += nums[j]
			if sum == k {

				count++
			}
		}
	}

	return count
}
