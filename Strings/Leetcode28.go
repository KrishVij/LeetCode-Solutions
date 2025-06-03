package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {

	if len(haystack) == 0 || len(needle) == 0 {

		return -1
	}

	Left,Right,Current,Length := 0,0,"",0
	
	for Right < len(haystack) {

		Current += (string)haystack[Right]
		Length = Right - Left + 1

		if Current == needle {

			return strings.IndexByte(current,needle[0])
		}

		if Length > len(needle) {

			Current = Current[left + 1 :]
			
		}

		Right++
	}

	return -1
	
}
