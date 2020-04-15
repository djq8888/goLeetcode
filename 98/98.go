package main

import "github.com/djq8888/goBinaryTree"

const (
	intMax = int(^uint(0)>>1)
	intMin = ^intMax
)

func checkVailidBST(root *goBinaryTree.TreeNode, lowBound, upbound int) bool {
	if root == nil {
		return true
	} else if root.Val >= upbound || root.Val <= lowBound {
		return false
	} else {
		return checkVailidBST(root.Left, lowBound, root.Val) && checkVailidBST(root.Right, root.Val, upbound)
	}
}

func isValidBST(root *goBinaryTree.TreeNode) bool {
	return checkVailidBST(root, intMin, intMax)
}
