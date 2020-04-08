package main

import (
	"fmt"
	"github.com/djq8888/goSingleList"
)

func main()  {
	nums1 := []int{2,4,3}
	nums2 := []int{5,6,4}
	list1 := goSingleList.Build(nums1)
	list2 := goSingleList.Build(nums2)
	sum := addTwoNumbers(list1, list2)
	fmt.Println(goSingleList.ToSlice(sum))
}
