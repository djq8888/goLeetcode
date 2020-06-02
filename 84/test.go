package main

import "fmt"

func main()  {
	//heights := []int{2,1,5,6,2,3}
	heights := []int{0,3,2,5}
	//heights := []int{1}
	//heights := []int{2,1,2}
	maxArea := largestRectangleArea(heights)
	fmt.Println(maxArea)
}