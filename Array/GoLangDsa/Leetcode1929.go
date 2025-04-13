package main

import "fmt"

func getConcatenation(nums []int) []int {

	length := 2 * len(nums)
	concatenatedArray := make([]int, length)

	for i := 0; i < length; i++ {

		concatenatedArray[i] = nums[i%len(nums)]
	}

	return concatenatedArray
}

func main() {

	nums := []int{1, 2, 1}
	result := getConcatenation(nums)

	fmt.Print(result)

}
