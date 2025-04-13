package main

func maximumUniqueSubarray(nums []int) int {

	if len(nums) == 0 {

		return 0
	}

	hash := make(map[int]bool)
	Left, Right, sum, maxSum := 0, 0, 0, 0

	for Right < len(nums) {

		for hash[nums[Right]] {

			sum -= nums[Left]
			delete(hash, nums[Left])
			Left++
		}

		hash[nums[Right]] = true
		sum += nums[Right]
		if sum > maxSum {

			maxSum = sum
		}
	}

	return maxSum
}
