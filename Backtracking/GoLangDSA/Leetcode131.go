package main

import (
	"fmt"
	"strings"
)

func chopFromLast(arr []string) []string {

	if len(arr) <= 0 {

		return []string{}
	}

	return arr[:len(arr)-1]
}

func pallindromeCheck(s string, start int, end int) bool {

	for start < end {

		if s[start] != s[end] {

			return false
		}

		start++
		end--
	}

	return true
}

func findAllPartitioning(s string, result *[][]string, substringStor []string, index int) {

	if index == len(s) {

		substringcpy := append([]string{}, substringStor...)
		*result = append(*result, substringcpy)

		return
	}

	for i := index; i < len(s); i++ {

		if pallindromeCheck(s, index, i) {

			substringStor = append(substringStor, s[index:i+1])
			findAllPartitioning(s, result, substringStor, i+1)
			substringStor = chopFromLast(substringStor)
		}
	}

}

func partition(s string) [][]string {

	var result [][]string

	findAllPartitioning(s, &result, []string{}, 0)

	return result
}

func main() {

	s := "aab"

	result := partition(s)

	for _, partition := range result {
		fmt.Println(strings.Join(partition, " | "))
	}
}
