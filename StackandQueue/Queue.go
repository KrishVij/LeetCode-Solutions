package main

import "fmt"

func Append(arr *[]int, x int) *[]int {

	*arr = append(*arr, x)

	return arr
}

func Remove(arr *[]int) {

	if len(*arr) == 0 {

		return
	}

	Left, Right := 0, len(*arr)-1

	*arr = (*arr)[Left+1:]
}

func printTop(arr []int) int {

	if len(arr) == 0 {
		return -1
	}

	return arr[0]
}

func Size(arr []int) int {

	Left, Right := 0, len(arr)-1

	return Right - Left + 1
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
