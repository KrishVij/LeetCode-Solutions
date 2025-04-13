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

func main() {

	var nums []int
	Push(&nums, 1)
	Push(&nums, 2)
	Remove(&nums)
	fmt.Println(printTop(nums))
	fmt.Println(Size(nums))
	fmt.Print(nums)
}
