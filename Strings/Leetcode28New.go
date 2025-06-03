package main

import (

	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {

	if len(haystack) == 0 || len(needle) == 0 {

		return -1
	}

	Right,nIndex := 0,0

	for Right < len(haystack) {

		if haystack[Right] == needle[nIndex] {

			nIndex++
		}else {

			Right = Right - nIndex

			nIndex = 0
		}

		if nIndex == len(needle) {

			return Right - len(needle) + 1
		}

		Right++
	}

	return -1
}
