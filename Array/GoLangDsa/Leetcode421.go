package main

func findMaximumXOR(nums []int) int {

	if len(nums) <= 0 {

		return 0
	}

	var xorStore []int

	for i := 0; i < len(nums); i++ {

		for j := i; j < len(nums); j++ {

			xor := nums[i] ^ nums[j]
			xorStore = append(xorStore, xor)
		}
	}

	max := xorStore[0]
	for k := 1; k < len(xorStore); k++ {

		if xorStore[k] > max {

			max = xorStore[k]
		}
	}

	return max

}
