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

func nextLessElement(nums []int) []int {

	if len(nums) == 0 {

		return []int{-1}
	}

	var stack []int
	Nle := make([]int,len(nums))

	for i := 0; i < len(nums);i++ {

		for Size(stack) != 0 && nums[i] > nums[getTop(stack)] {

			Nle[i] = nums[getTop(stack)]
			Remove(&stack)
		}

		Push(&stack,i)
	}
}
