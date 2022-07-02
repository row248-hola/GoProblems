package main

import (
	"testing"
)

func TestValidParentheses1(t *testing.T) {
	number := 1

	expected := map[string]bool{
		"()": false,
	}

	result := GenerateParenthesis(number)

	for i := 0; i < len(result); i++ {
		if _, ok := expected[result[i]]; !ok {
			t.Errorf("Invalid parentheses %s", result[i])
			return
		}
		expected[result[i]] = true
	}

	var unmatchedValues []string
	for k, v := range expected {
		if v == false {
			unmatchedValues = append(unmatchedValues, k)
		}
	}

	if len(unmatchedValues) > 0 {
		t.Errorf("Do not catched values: %s", unmatchedValues)
	}
}

func TestValidParentheses2(t *testing.T) {
	number := 3

	expected := map[string]bool{
		"((()))": false,
		"(()())": false,
		"(())()": false,
		"()(())": false,
		"()()()": false,
	}

	result := GenerateParenthesis(number)

	for i := 0; i < len(result); i++ {
		if _, ok := expected[result[i]]; !ok {
			t.Errorf("Invalid parentheses %s", result[i])
			return
		}
		expected[result[i]] = true
	}

	var unmatchedValues []string
	for k, v := range expected {
		if v == false {
			unmatchedValues = append(unmatchedValues, k)
		}
	}

	if len(unmatchedValues) > 0 {
		t.Errorf("Do not catched values: %s", unmatchedValues)
	}
}
