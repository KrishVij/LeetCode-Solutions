package main

import "fmt"

func calculateNthFibonacci(n int) int {

	if n == 0 {

		return 0
	}

	if n == 1 {

		return 1
	}

	return calculateNthFibonacci(n-1) + calculateNthFibonacci(n-2)
}

func fib(n int) int {

	return calculateNthFibonacci(n)
}

func main() {

	n := 2

	result := fib(n)

	fmt.Print(result)
}
