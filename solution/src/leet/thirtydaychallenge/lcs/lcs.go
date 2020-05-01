package main

import "fmt"

var maxVal = 1000000000000000000

func lcs(i int, j int, text1 string, text2 string, dp [][]int, N int, M int) int {
	if i >= N || j >= M {
		return 0
	}
	if dp[i][j] != maxVal {
		return dp[i][j]
	}
	if text1[i] == text2[j] {
		return 1 + lcs(i+1, j+1, text1, text2, dp, N, M)
	}

	n := lcs(i+1, j, text1, text2, dp, N, M)
	m := lcs(i, j+1, text1, text2, dp, N, M)

	dp[i][j] = getMax(n, m)

	return dp[i][j]

}

func getMax(n int, m int) int {
	if m > n {
		return m
	}
	return n
}

func longestCommonSubsequence(text1 string, text2 string) int {
	nLen := len(text1)
	mLen := len(text2)

	dp := make([][]int, nLen)
	for i := range dp {
		dp[i] = make([]int, mLen)
		for j := 0; j < mLen; j++ {
			dp[i][j] = maxVal
		}

	}

	return lcs(0, 0, text1, text2, dp, nLen, mLen)
}

func main() {
	str1 := "ABCDGH"
	str2 := "AEDFHR"
	fmt.Println(longestCommonSubsequence(str1, str2))
}
