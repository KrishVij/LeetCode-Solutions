package main

import (
	"fmt"
	"math"
)

func powerCalculate(n int, x int) int {

	if x == 0 {

		return 1
	}

	if x%2 == 0 {

		half := powerCalculate(n, x/2)
		return half * half
	} else {

		return n * powerCalculate(n, x-1)
	}
}

func checkPower(n, x int) bool {

	if powerCalculate(2, x) > n {

		return false
	}

	if powerCalculate(2, x) == n {

		return true
	}
	return checkPower(n, x+1)
}
func isPowerofTwo(n int) bool {

	return checkPower(n, 0)
}
