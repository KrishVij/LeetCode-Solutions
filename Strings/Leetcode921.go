package main

func minAddToMakeValid(s string) int {

	openCount, closeCount := 0, 0

	for index := 0; index <= len(s)-1; index++ {

		if s[index] == '(' {

			openCount++
		} else if s[index] == ')' {

			if openCount > 0 {

				openCount--
			} else {

				closeCount++
			}

		}
	}

	return openCount + closeCount
}
