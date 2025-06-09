package main

import (
	"strings"
	"unicode"
)

func decodeString(s string) string {

	numStack := []int{}
	alphaStack := []string{}
	
	currentNum := 0
	Result := ""

	for _,Char := range s {

		if unicode.IsDigit(Char) {

			currentNum = currentNum * 10 + int(Char - '0')
		}else if Char == '[' {

			numStack = append(numStack, currentNum)
			alphaStack = append(alphaStack, Result)

			currentNum = 0
			Result = ""
		}else if Char == ']' {

			previousNum := numStack[len(numStack) - 1]
			numStack = numStack[:len(numStack) - 1]

			previousString := alphaStack[len(alphaStack) - 1]
			alphaStack = alphaStack[:len(alphaStack) - 1]

			Result = previousString + strings.Repeat(Result,previousNum)
		}else {

			Result += string(Char)
		}
		
	} 

	return Result
}
