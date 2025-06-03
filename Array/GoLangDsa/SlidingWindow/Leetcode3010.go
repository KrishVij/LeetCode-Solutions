package main

func minimumCost(nums []int) int {

	if len(nums) == 0 {

		return 0
	}

	Left, Right := 0, 0
	Count := 0
	Sum := 0

	for Right < len(nums) {

		Sum += nums[Right]
		Count++
		if Right - Left == 1 || Count > 3 {

			Left++
		}

		Right++
	}

	return Sum
}