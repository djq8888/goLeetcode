package main

import (
	"fmt"
	"github.com/djq8888/goBinaryTree"
)

func main()  {
	//num := []int{5,1,4,-1,-1,3,6}
	num := []int{2,1,3}
	root := goBinaryTree.BuildLevelOrder(num)
	fmt.Println(isValidBST(root))
}
