package main

import (

	"fmt"
	"strconv"
)

func decodeString(s string) string {

	if len(s) == 0 {

		return ""
	}

	Str := ""
	decodedstring := ""
	for i := 0;i < len(s);i++ {

		
		if s[i] >= '0' && s[i] <= '9' {

			k,_ := strconv.Atoi(s[i])
		}else if s[i] >= 'a' && s[i] <= 'z' {

			
			Str += s[i]
		}else if s[i] == ']' {

			decodedString += (Str * k)
			Str = ""
		}
	}

	return decodedString
}
