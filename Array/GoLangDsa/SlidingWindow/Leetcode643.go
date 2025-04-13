package main

func findMaxAverage(nums []int, k int) float64 {

	if len(nums) == 0 {

		return 0.0
	}

	Left := 0
	Right := 0
	Average := 0.0
	Length := 0
	maxAverage := -1e9

	for Right < len(nums) {

		Length = Right - Left + 1

		Average = Average + float64(nums[Right])

		if Length > k {

			Average = Average - float64(nums[Left])
			Left++
			Length--
		}
		if Length == k {

			currentAverage := Average / float64(Length)
			maxAverage = max(maxAverage, currentAverage)
		}

		Right++
	}

	return maxAverage

}
