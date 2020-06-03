package main

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1)==0 || len(text2)==0 {
		return 0
	}

	//创建二维数组
	dp := make([][]int, len(text1)+1, len(text1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1, len(text2)+1)
	}
	
	//动态规划
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			plus := 0
			if text1[i-1] == text2[j-1] {
				plus = 1
			}
			dp[i][j] = max3(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]+plus)
		}
	}

	return dp[len(text1)][len(text2)]
}

func max3(a, b, c int) int {
	if b <= a && c <= a {
		return a
	} else if a <= b && c <= b {
		return b
	} else {
		return c
	}
}