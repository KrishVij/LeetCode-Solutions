package main

import "fmt"

func CalculateSubsetSum(nums []int, Targetsum *[]int, sum int, index int) {

	if index >= len(nums) {

		*Targetsum = append(*Targetsum, sum) //we are only appending the sum when base case is beig reached there is hno point to store in betweem
		return
	}

	CalculateSubsetSum(nums, Targetsum, sum+nums[index], index+1)

	CalculateSubsetSum(nums, Targetsum, sum, index+1)
}

func subsetSum(nums []int) []int {

	var result []int

	CalculateSubsetSum(nums, &result, 0, 0)

	return result

}

func main() {

	nums := []int{5, 2, 1}
	fmt.Println(subsetSum(nums))
}
