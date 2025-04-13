package main

import "fmt"

func chopFromLast(arr []int) []int {

	if len(arr) <= 0 {

		return []int{}
	}

	return arr[:len(arr)-1]
}

func findAllPermutation(nums []int, permutationStor []int, check map[int]bool, result *[][]int) {

	if len(permutationStor) == len(nums) {

		subsetcpy := append([]int{}, permutationStor...)
		*result = append(*result, subsetcpy)
		return
	}

	for i := 0; i < len(nums); i++ {

		if !check[nums[i]] {

			permutationStor = append(permutationStor, nums[i])
			check[nums[i]] = true
			findAllPermutation(nums, permutationStor, check, result)
			check[nums[i]] = false
			permutationStor = chopFromLast(permutationStor)

		}
	}
}

func permute(nums []int) [][]int {

	var result [][]int
	check := make(map[int]bool)

	findAllPermutation(nums, []int{}, check, &result)

	return result
}

func main() {

	nums := []int{1, 2, 3}
	result := permute(nums)

	for _, val := range result {

		fmt.Println(val)
	}
}
