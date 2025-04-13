package main

import "fmt"

func findKthLargest(nums []int, k int) int {

	for i := 0; i < len(nums); i++ {

		largestIdx := 0
		for j := 1; j < len(nums); j++ {

			if nums[j] > nums[largestIdx] {

				largestIdx = j
			}
		}

		if i == k-1 {

			return nums[largestIdx]
		}

		nums = append(nums[:largestIdx], nums[largestIdx+1:]...)
	}

	return -1

}
