package main

func circularArrayLoop(nums []int) bool {

	if len(nums) < 2 {

		return false
	}

	indicesStore := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if nums[i] == 0 {
			continue
		}

		j := i

		direction := nums[j] > 0

		for {

			if _, exists := indicesStore[j]; exists {

				return true
			}

			indicesStore[j] = nums[j]

			next := (j + nums[j]) % len(nums)

			if next < 0 {
				next += len(nums)
			}

			if next == j || (nums[next] > 0) != direction {

				break
			}

			j = next

		}

	}

	return false
}
