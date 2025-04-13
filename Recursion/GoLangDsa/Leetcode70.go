package main

import "fmt"

func toReachtheTop(n, bottom int) int {

	if bottom == n {

		return 0
	}

	if bottom > n {

		return 0
	}

	return toReachtheTop(n, bottom+1) + toReachtheTop(n, bottom+2)
}

func climbingStairs(n int) int {

	return toReachtheTop(n, 0)

}
