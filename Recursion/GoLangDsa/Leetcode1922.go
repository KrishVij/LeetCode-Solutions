package main

import "fmt"

const MOD = 1_000_000_007

func recursiveCount(remaining, index int) int64 {

	if remaining == 0 {

		return 1
	}

	choices := 5
	if index%2 == 1 {

		choices = 4
	}

	return (int64(choices) * recursiveCount(remaining-1, index+1)) % MOD
}

func countGoodNumbers(n int64) int {

	return int(recursiveCount(int(n), 0) % MOD)
}
