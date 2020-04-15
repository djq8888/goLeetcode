package main

import "github.com/djq8888/goBinaryTree"

func mvLeftToRight(left, right *goBinaryTree.TreeNode) *goBinaryTree.TreeNode {
	if right == nil {
		return left
	}
	insert := right
	for insert.Left != nil {
		insert = insert.Left
	}
	insert.Left = left
	return right
}

func deleteNode(root *goBinaryTree.TreeNode, key int) *goBinaryTree.TreeNode {
	if root == nil {
		return nil
	} else if root.Val == key {
		return mvLeftToRight(root.Left, root.Right)
	} else {
		if key > root.Val {
			root.Right = deleteNode(root.Right, key)
		} else {
			root.Left = deleteNode(root.Left, key)
		}
	}
	return root
}
