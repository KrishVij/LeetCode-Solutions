package main

import "fmt"

func Push(arr *[]int, x int) {

	*arr = append(*arr, x)

}

func Remove(arr *[]int) {

	if len(*arr) == 0 {

		return
	}

	*arr = (*arr)[:len(*arr)-1]
}

func printTop(arr []int) int {

	if len(arr) == 0 {

		return -1
	}
	return arr[len(arr)-1]
}

func Size(arr []int) int {

	if len(arr) == 0 {

		return 0
	}

	return len(arr)
}

func previousLessElement(nums []int) []int {

	if len(nums) == 0 {

		return []int{-1}
	}

	var stack []int
	Ple := make([]int,len(nums))

	for i := 0;i < len(nums);i++ {

		for Size(stack) != 0 && nums[i] < arr[getTop(stack)] {

			Remove(&stack)
		}

		if Size(stack) == 0 {

			Ple[i] == -1
		}else {

			Ple[i] = Ple[getTop(stack)]
		}

		Push(&stack,i)
	}
}
