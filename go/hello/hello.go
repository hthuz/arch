package main

import (
	"fmt"
	"math"
)

func main() {
	// value := [128]bool{}
	// fmt.Println(value)
	ans := coinChange([]int{2}, 3)
	fmt.Println(ans)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] == math.MaxInt32 {
		fmt.Println("here")
		return -1
	} else {
		return dp[amount]
	}

}

func longestCommonSubsequence(text1 string, text2 string) int {
	width := len(text1)
	height := len(text2)
	dp := make([][]int, height)
	for i := range height {
		dp[i] = make([]int, width)
	}
	for i := range height {
		if text1[0] == text2[i] {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	for i := range width {
		if text1[i] == text2[0] {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}
	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			if text1[j] == text2[i] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	for _, row := range dp {
		fmt.Println(row)
	}
	return dp[height-1][width-1]
}
