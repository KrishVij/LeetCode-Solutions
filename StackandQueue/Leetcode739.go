package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {

	if len(temperatures) == 0 {

		return []int{0}
	}

	stack := []int{}
	waitForWarmerDays := make([]int,len(temperatures))

	for i := 0,i < len(temperatures);i++ {

		for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < temperatures[i] {

			topIndex := stack[len(stack) - 1]
			stack = [:len(satck) - 1]

			waitForWarmerDays[topIndex] = i - topIndex
		}

		stack = append(stack,i)
	}

	return waitForWarmerDays
}
