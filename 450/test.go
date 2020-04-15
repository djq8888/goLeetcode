package main

import (
	"fmt"
	"github.com/djq8888/goBinaryTree"
)

func main()  {
	root := goBinaryTree.BuildLevelOrder([]int{5,3,6,2,4,-1,7})
	key := 3
	root = deleteNode(root, key)
	fmt.Println("add breakpoint hear.")
}
