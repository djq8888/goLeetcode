package main

import (
	"github.com/djq8888/goBinaryTree"
)

func insert(root *goBinaryTree.TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = new(goBinaryTree.TreeNode)
			root.Left.Val = val
			return
		}
		insert(root.Left, val)
	} else {
		if root.Right == nil {
			root.Right = new(goBinaryTree.TreeNode)
			root.Right.Val = val
			return
		}
		insert(root.Right, val)
	}
}

func insertIntoBST(root *goBinaryTree.TreeNode, val int) *goBinaryTree.TreeNode {
	if root == nil {
		return nil
	}
	insert(root, val)
	return root
}
