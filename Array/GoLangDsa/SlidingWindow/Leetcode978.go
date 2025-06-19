package main

func compare (a,b int) int {

	if a < b {

		return -1
	}else if a > b {

		return 1
	}

	return 0
}

func max (a,b int) int {

	if a > b {

		return a
	}

	return b
}

func maxTurbulenceSize(arr []int) int{

	length := 1
	maxLen := 1
	prevCmp := 0

	for i := 1;i < len(arr);i++ {

		currCmp := compare(arr[i - 1],arr[i])

		if currCmp == 0 {

			length = 1
		}else if prevCmp * currCmp == -1 {

			length++
		}else {

			length = 2
		}

		maxLen = max(maxLen,length)
		prevCmp = currCmp
	}

	return maxLen
}
