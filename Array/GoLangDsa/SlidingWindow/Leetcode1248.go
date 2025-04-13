package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {

	Left := 0
	Right := 0
	Length := 0
	maxLen := 0

	for Right < len(nums) {

		if Right-Left+1 == 1 {

			Right++
			continue
		}

		l, r := Left, Right
		andResult := nums[l]
		storeAnd := []int{}
		for l <= r {

			andResult &= nums[l]
			storeAnd = append(storeAnd, andResult)
			l++
		}

		sum := 0
		for i := 0; i < len(storeAnd); i++ {

			sum += storeAnd[i]
		}

		if sum == 0 {

			Length = Right - Left + 1
			if Length > maxLen {
				maxLen = Length
			}

		} else {

			Left++
		}

		Right++

	}

	return maxLen

}
