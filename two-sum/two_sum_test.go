package main

import (
	"encoding/json"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expectedIndexSum := 1

	result := TwoSumHashTableImplementation(nums, target)
	if result[0]+result[1] != expectedIndexSum {
		resultToString, _ := json.Marshal(result)
		t.Errorf("Got %s expected sum %v", resultToString, expectedIndexSum)
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expectedIndexSum := 1

	result := TwoSumHashTableImplementation(nums, target)
	if result[0]+result[1] != expectedIndexSum {
		resultToString, _ := json.Marshal(result)
		t.Errorf("Got %s expected sum %v", resultToString, expectedIndexSum)
	}
}

func TestTwoSum3(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expectedIndexSum := 1

	result := TwoSumHashTableImplementation(nums, target)
	if result[0]+result[1] != expectedIndexSum {
		resultToString, _ := json.Marshal(result)
		t.Errorf("Got %s expected sum %v", resultToString, expectedIndexSum)
	}
}

func TestTwoSum4(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expectedIndexSum := 3

	result := TwoSumHashTableImplementation(nums, target)
	if result[0]+result[1] != expectedIndexSum {
		resultToString, _ := json.Marshal(result)
		t.Errorf("Got %s expected sum %v", resultToString, expectedIndexSum)
	}
}
