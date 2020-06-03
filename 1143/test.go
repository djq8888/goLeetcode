package main

import "fmt"

func main()  {
	text1 := "abcde"
	text2 := "ace"
	l := longestCommonSubsequence(text1, text2)
	fmt.Println(l)
}
