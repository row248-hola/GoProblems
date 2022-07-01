package main

import "github.com/emirpasic/gods/stacks/linkedliststack"

func isCloseTag(tagToClose string, tag string) bool {
	switch tagToClose {
	case "(":
		return tag == ")"
	case "[":
		return tag == "]"
	case "{":
		return tag == "}"
	}

	return false
}

func isOpenTag(tag string) bool {
	return tag == "(" || tag == "[" || tag == "{"
}

func IsValidParentheses(s string) bool {
	stack := linkedliststack.New()

	for i := 0; i < len(s); i++ {
		if isOpenTag(string(s[i])) {
			stack.Push(string(s[i]))
			continue
		}

		if stack.Size() == 0 {
			return false
		}

		tagToClose, _ := stack.Pop()
		if isCloseTag(tagToClose.(string), string(s[i])) {
			continue
		}

		return false
	}

	return stack.Size() == 0
}
