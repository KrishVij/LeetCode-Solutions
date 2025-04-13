package main

func Push(arr *[]byte, x byte) {

	*arr = append(*arr, x)

}

func Remove(arr *[]byte) {

	if len(*arr) == 0 {

		return
	}

	*arr = (*arr)[:len(*arr)-1]
}

func printTop(arr []byte) int {

	if len(arr) == 0 {

		return -1
	}
	return arr[len(arr)-1]
}

func Size(arr []byte) int {

	if len(arr) == 0 {

		return 0
	}

	return len(arr)
}

func isValid(s string) bool {

	if len(s) == 0 {

		return false
	}

	var stack []byte

	for i := 0; i < len(s); i++ {

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {

			Push(&stack, s[i])
		} else {

			if len(stack) == 0 {
				return false
			}
		}

		ch := printTop(stack)
		Remove(&stack)
		if s[i] == ')' && ch == '(' || s[i] == ']' && ch == '[' || s[i] == '}' && ch == '{' {

			continue
		} else {

			return false
		}
	}

	return isEmpty(stack)
}
