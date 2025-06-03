package main

import "fmt"

func countSubarrays(nums []int) int {

	if len(nums) == 0 {

		return 0
	}

	Left,Right,Count := 0,0,0

	for Right < len(nums) {

		Length := Right - Left + 1

		if Length == 3 {

			if  float64(nums[Left] + nums[Right]) ==(float64(nums[Left + 1]) / 2) {

				Count++
			}
		}
		
		if Length > 3 {

			Left++
		}

		Right++
		
	}

	return Count

	
}
