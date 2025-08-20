package main

func findTheDifference(s string, t string) byte {

	  if len(t) != len(s) + 1 {
        return 0
    }

	for i := 0; i < len(s); i++ {

		if s[i] != t[i] {

			return t[i]
		}
	}

	return t[len(t) - 1]
}