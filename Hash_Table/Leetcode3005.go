package main

func maxFrequencyElements(nums []int) int {

	if len(nums) == 0 {

		return 0
	}

	frequency := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		frequency[nums[i]]++
	}

	maxFreq := 0

	for _, val := range frequency {

		if val > maxFreq {

			maxFreq = val
		}
	}

	sum := 0
	for _, val := range frequency {

		if val == maxFreq {

			sum += maxFreq
		}
	}

	return sum
}
