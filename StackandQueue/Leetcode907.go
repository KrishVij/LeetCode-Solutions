const MOD = 1_000_000_007

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

func sumSubarrayMins(arr []int) int {

	if len(arr) == 0 {

		return -1
	}

	var stack []int
	Ple := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {

		for Size(stack) != 0 && arr[printTop(stack)] >= arr[i] {

			Remove(&stack)
		}

		if Size(stack) == 0 {

			Ple[i] = -1
		} else {

			Ple[i] = printTop(stack)
		}

		Push(&stack, i)
	}

	Nle := make([]int, len(arr))
	stack = []int{}

	for i := len(arr) - 1; i >= 0; i-- {

		for Size(stack) != 0 && arr[printTop(stack)] >arr[i] {

			Remove(&stack)
		}

		if Size(stack) == 0 {

			Nle[i] = len(arr)
		} else {

			Nle[i] = printTop(stack)
		}

		Push(&stack, i)
	}

	Sum := 0
	for i := 0; i < len(arr); i++ {

		Left := i - Ple[i]
		Right := Nle[i] - i

		Sum = Sum + (arr[i]*Left*Right)%MOD
	}

	return Sum % MOD
}
