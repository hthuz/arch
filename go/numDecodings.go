package main

import "fmt"

func main() {
	numDecodings("2101")
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	if len(s) == 1 {
		return dp[0]
	}

	if (s[0]-'0')*10+(s[1]-'0') <= 26 {
		dp[1] += 1
	}

	if s[1] != '0' {
		dp[1] += 1
	}

	for i := 2; i < len(s); i++ {
		if (s[i-1]-'0')*10+(s[i]-'0') <= 26 && s[i-1] != '0' {
			dp[i] += dp[i-2]
		}
		if s[i] != '0' {
			dp[i] += dp[i-1]
		}
	}
	fmt.Println(dp)

	return dp[len(dp)-1]
}
