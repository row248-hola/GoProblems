package main

import (
	"testing"
)

func TestValidParentheses1(t *testing.T) {
	str := "()"

	result := IsValidParentheses(str)
	if result != true {
		t.Errorf("Got %t expected %t", result, true)
	}
}

func TestValidParentheses2(t *testing.T) {
	str := "()[]{}"

	result := IsValidParentheses(str)
	if result != true {
		t.Errorf("Got %t expected %t", result, true)
	}
}

func TestValidParentheses3(t *testing.T) {
	str := "(]"

	result := IsValidParentheses(str)
	if result != false {
		t.Errorf("Got %t expected %t", result, false)
	}
}

func TestValidParentheses4(t *testing.T) {
	str := "({)}"

	result := IsValidParentheses(str)
	if result != false {
		t.Errorf("Got %t expected %t", result, false)
	}
}
