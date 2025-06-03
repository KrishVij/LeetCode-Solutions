package main

import (

	"fmt"
	"strconv"
)

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

func calPoints(operations []string) int {

	if len(operations) == 0 {

		return 0
	}

	var stack []int
    var otherStack []int
	Sum := 0
    Ans := 0

	for i := 0; i < len(operations); i++ {

		if operations[i] == "D" {

			score := 2 * getTop(stack)
			Push(&stack, score)
		} else if operations[i] == "C" {

			Remove(&stack)
		} else if operations[i] == "+" {

            Top := getTop(stack)
            Remove(&stack)
            newTop := getTop(stack)
            Remove(&stack)
            Total := Top + newTop
            Push(&stack,newTop)
            Push(&stack,Top)
            Push(&stack,Total)

		} else {
			x, _ := strconv.Atoi(operations[i])
			Push(&stack, x)
		}
	}

    for Size2(stack) != 0 {

        Sum += getTop(stack)
        Remove(&stack)
    }

	return Sum
}
