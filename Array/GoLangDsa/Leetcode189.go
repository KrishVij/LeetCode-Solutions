package main

func reverseArray(nums []int, begin, end int) {

	for begin < end {

		temp := nums[begin]
		nums[begin] = nums[end]
		nums[end] = temp

		begin++
		end--
	}

}

func rotate(nums []int, k int) {

	if len(nums) == 0 {

		return
	}

	reverseArray(&nums, 0, len(nums)-1)
	reverseArray(&nums, 0, k-1)
	reverseArray(&nums, len(nums)-k, len(nums)-1)
}
