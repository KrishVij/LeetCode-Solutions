package main

import "slices"

func divideArray(nums []int, k int) [][]int {

	if k == 0 || len(nums) == 0 {

		return [][]int{}
	}

	slices.Sort(nums)

	divisionArray := make([][]int,len(nums)/3)

	for i := 0; i < len(nums); i+=3 {

        if nums[i+2] - nums[i] <= k {

            coloumnData := make([]int,3)
			coloumnData[0] = nums[i]
			coloumnData[1] = nums[i+1]
			coloumnData[2] = nums[i+2]

			divisionArray[i/3] = coloumnData
        }else {

            return [][]int{}
        }
	}

    return divisionArray

}