package main

import (
	"github.com/djq8888/goBinaryTree"
)

func inOrder(root *goBinaryTree.TreeNode, array *[]*goBinaryTree.TreeNode)  {
	if root == nil {
		return
	}
	inOrder(root.Left, array)
	*array = append(*array, root)
	inOrder(root.Right, array)
}

func sortedArrayToBst(arr []*goBinaryTree.TreeNode) *goBinaryTree.TreeNode {
	if len(arr) == 0 {
		return nil
	} else {
		mid := len(arr) / 2
		root := arr[mid]
		root.Left = sortedArrayToBst(arr[:mid])
		root.Right = sortedArrayToBst(arr[mid+1:])
		return root
	}
}

func balanceBST(root *goBinaryTree.TreeNode) *goBinaryTree.TreeNode {
	//先中序遍历得到有序数组
	array := make([]*goBinaryTree.TreeNode, 0, 10000)
	inOrder(root, &array)

	//然后构造平衡二叉搜索树
	return sortedArrayToBst(array)
}