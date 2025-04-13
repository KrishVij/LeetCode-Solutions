package main

import "fmt"

func findPeakElement(nums []int) int {

	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {

			return i
		}
	}

	return -1
}
