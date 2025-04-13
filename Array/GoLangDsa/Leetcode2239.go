func abs(x int) int {

	if x < 0 {

		return -x
	} else {

		return x
	}
}

func findClosestNumber(nums []int) int {

	closest := nums[0]

	for _, num := range nums {

		if abs(num) < abs(closest) || abs(num) == abs(closest) && num > closest {

			closest = num
		}
	}

	return closest
}
