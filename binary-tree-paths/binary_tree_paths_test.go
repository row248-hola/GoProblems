package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"testing"
)

func TestBinaryTreePaths1(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
	}

	expectedPaths := hashset.New()
	expectedPaths.Add("1")

	paths := BinaryTreePaths(tree)

	for i := 0; i < len(paths); i++ {
		if !expectedPaths.Contains(paths[i]) {
			t.Errorf("Invalid path %s", paths[i])
		}

		expectedPaths.Remove(paths[i])
	}

	if expectedPaths.Size() > 0 {
		json, _ := expectedPaths.ToJSON()
		t.Errorf("Paths %s was not found", string(json))
	}
}

func TestBinaryTreePaths2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	expectedPaths := hashset.New()
	expectedPaths.Add("1->2->5", "1->3")

	paths := BinaryTreePaths(tree)

	for i := 0; i < len(paths); i++ {
		if !expectedPaths.Contains(paths[i]) {
			t.Errorf("Invalid path %s", paths[i])
		}

		expectedPaths.Remove(paths[i])
	}

	if expectedPaths.Size() > 0 {
		json, _ := expectedPaths.ToJSON()
		t.Errorf("Paths %s was not found", string(json))
	}
}
