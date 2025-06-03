package Strings

import (
	"fmt"
)

// func isPallindrome(s string) bool {

// 	if len(s) == 0 {

// 		return false
// 	}

// 	str := ""

// 	for i := 0;i < len(s);i++ {

// 		str += s[i]
// 	}

// 	return str == s
// }
// func longestPalindrome(s string) string {

// 	if len(s) == 0 {

// 		return ""
// 	}

// 	subStrPallindrome := ""
	

// 	for i := 0;i < len(s);i++ {

// 		current := ""
// 		for j := i;j < len(s);j++ {

// 			current += s[i]
			
// 			if isPallindrome(current) && len(current) > maxLen {

// 				subStrPallindrome = current
// 				maxLen = len(current)
// 			}
// 		}
// 	}

// 	return subStrPallindrome
//}

func checkPallindrome(s string,left,right int) bool {

	if len(s) == 0 {

		return ""
	}

	for left < right {

		if s[left] != s[right] {

			return false
		}

		left++
		right--
	}

	return true
}
