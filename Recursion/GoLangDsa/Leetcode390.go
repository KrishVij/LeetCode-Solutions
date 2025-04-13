package main

func findLastElement(numbers []int, leftToRight bool) int {

	if len(numbers) == 1 {

		return numbers[len(numbers)]
	}

	var newNumbers []int

	if leftToRight {

		for i := 0; i < len(numbers); i += 2 {

			newNumbers = append(newNumbers, numbers[i])
		}
	} else {

		for i := len(numbers) % 2; i < len(numbers); i += 2 {

			newNumbers = append(newNumbers, numbers[i])
		}
	}

	return findLastElement(numbers, !leftToRight)
}

func lastRemaining(n int) int {

	var numbers []int

	for i := 1; i <= n; i++ {

		numbers[i] = i
	}

	return findLastElement(numbers, true)

}
