package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(path string, node *TreeNode) {
	if node.Left != nil {
		traverse(path+"->"+strconv.Itoa(node.Left.Val), node.Left)
	}

	if node.Right != nil {
		traverse(path+"->"+strconv.Itoa(node.Right.Val), node.Right)
	}

	if node.Left == nil && node.Right == nil {
		result = append(result, path)
	}
}

var result []string

func BinaryTreePaths(root *TreeNode) []string {
	result = []string{}

	traverse(strconv.Itoa(root.Val), root)

	return result
}
