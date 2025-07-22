package main

func singleNumber(nums []int) int {

	if len(nums) <= 0 {

		return -1
	}

	mpp := make(map[int]int)

	for _,val := range nums {

		mpp[val]++
	}

	for key,val := range mpp {

		if val < 3 {

			return key
		}
	}

	return -1
}