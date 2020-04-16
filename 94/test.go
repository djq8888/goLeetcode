package main

import (
	"fmt"
	"github.com/djq8888/goBinaryTree"
)

func main()  {
	root := goBinaryTree.BuildLevelOrder([]int{1,-1,2,3})
	arr := inorderTraversal(root)
	fmt.Println(arr)
}
