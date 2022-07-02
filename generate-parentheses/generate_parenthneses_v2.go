package main

func isValid(str string) bool {
	var opened int
	for i := 0; i < len(str); i++ {
		if string(str[i]) == ")" {
			opened--
		} else {
			opened++
		}

		if opened < 0 {
			return false
		}
	}

	return opened == 0
}

func generate(str string, chars int) {
	if len(str) == chars {
		if isValid(str) {
			result = append(result, str)
		}
		return
	}

	generate(str+"(", chars)
	generate(str+")", chars)
}

var result []string

func GenerateParenthesisV2(n int) []string {
	chars := n * 2
	result = []string{}

	generate("(", chars)
	generate(")", chars)

	return result
}
