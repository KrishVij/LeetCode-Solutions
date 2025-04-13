package main

func isSubsequence(s string, t string) bool {

	if s == "" {

		return false
	}

	subsequence := ""
	Right := 0
	Left := 0

	for Right < len(t) {

		if s[Left] == t[Right] {

			subsequence += string(s[Left])
			Left++
		}

		Right++
	}

	return s == subsequence

}
