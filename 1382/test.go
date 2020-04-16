package main

import (
	"fmt"
	"github.com/djq8888/goBinaryTree"
	"github.com/pkg/profile"
)

func main()  {
	//开始性能分析, 返回一个停止接口
	stopper := profile.Start(profile.MemProfile, profile.ProfilePath("."))
	//在main()结束时停止性能分析
	defer stopper.Stop()
	//for i := 0; i < 100000; i++ {
		root := goBinaryTree.BuildLevelOrder([]int{1,-1,2,-1,3,-1,4,-1,-1})
		root = balanceBST(root)
	//}
	fmt.Println("add breakpoint here.")
}
