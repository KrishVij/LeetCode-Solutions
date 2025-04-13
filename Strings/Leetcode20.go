package main

func isValid(s string) bool {

	if len(s) == 0 {

		return false
	}

	var Brackets []byte

	for i := 0; i < len(s); i++ {

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {

			Brackets = append(Brackets, s[i])

		} else {
			if len(Brackets) == 0 {
				return false
			}

			j := len(Brackets) - 1

			if (Brackets[j] == '(' && s[i] == ')') || (Brackets[j] == '{' && s[i] == '}') || (Brackets[j] == '[' && s[i] == ']') {

				Brackets = Brackets[:j]

			} else {

				return false
			}
		}
	}

	return len(Brackets) == 0

}
