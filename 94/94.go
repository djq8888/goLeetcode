package main

import "github.com/djq8888/goBinaryTree"

func inorderTraversal(root *goBinaryTree.TreeNode) []int {
	if root == nil {
		return nil
	}
	r1 := inorderTraversal(root.Left)
	r2 := inorderTraversal(root.Right)
	return append(append(r1, root.Val), r2...)
}
