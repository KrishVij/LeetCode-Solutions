package main

import "fmt"

func swapLast2Digits(num1, num2 *int) {

	if num1 == num2 {

		return
	}

	temp := *num1
	*num1 = *num2
	*num2 = temp
}

func reverseArray(arr []int, k int) []int {

	left, right := k, len(arr)-1

	for left < right {

		swapLast2Digits(&arr[left], &arr[right])
		left++
		right--
	}

	return arr
}

func nextPermutation(nums []int) {

	n := len(nums)

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {

		i--
	}

	if i >= 0 {

		j := n - 1
		for nums[j] <= nums[i] {

			j--
		}

		swapLast2Digits(&nums[i], &nums[j])
	}

	reverseArray(nums, i+1)
}

func main() {

	nums := []int{1, 2, 3}
	nextPermutation(nums)

	for _, val := range nums {

		fmt.Print(val)
	}

}
