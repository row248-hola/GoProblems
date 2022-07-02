package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	OPEN  = "1" // "("
	CLOSE = "0" // ")"
)

func checkIsValid(str string) bool {
	var opened int
	for i := 0; i < len(str); i++ {
		if string(str[i]) == CLOSE {
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

func GenerateParenthesis(n int) []string {
	symbols := n * 2

	binaryStr := strings.Repeat(OPEN, symbols)

	binary, err := strconv.ParseInt(binaryStr, 2, 32)
	if err != nil {
		fmt.Println("can't parse ", binary, " to number")
		return nil
	}

	var result []string
	for i := 0; i < int(binary); i++ {
		str := fmt.Sprintf("%b", i)
		str = str + strings.Repeat(OPEN, len(binaryStr)-len(str))

		if checkIsValid(str) {
			str = strings.Replace(str, OPEN, "(", -1)
			str = strings.Replace(str, CLOSE, ")", -1)

			result = append(result, str)
		}
	}

	return result
}
