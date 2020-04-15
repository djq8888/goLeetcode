package main

import (
	"fmt"
	"github.com/djq8888/goBinaryTree"
)

func main()  {
	root := goBinaryTree.BuildLevelOrder([]int{4,2,7,1,3})
	val := 5
	root = insertIntoBST(root, val)
	fmt.Println("add breakpoint here.")
}
