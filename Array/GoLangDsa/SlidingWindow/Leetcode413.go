package main

func numberOfArithmeticSlices(nums []int) int {

	if len(nums) < 3 {

		return 0
	}

	Left := 0
	Right := 2
	Count := 0

	for Right < len(nums) {

		if nums[Right]-nums[Right-1] == nums[Right-1]-nums[Right-2] {

			Count += (Right - Left) - 1
		} else {

			Left = Right - 1
		}
		Right++
	}

	return Count

}
