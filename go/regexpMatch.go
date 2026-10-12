package main

import "fmt"

func main() {
	// s := "aabbccc"
	// p := "aab*c*"
	s := "aa"
	p := "a*"
	ans := isMatch(s, p)
	fmt.Println(ans)

}

func isMatch(s string, p string) bool {

	m, n := len(s), len(p)

	// s[:]
	dp := make([][]bool, m+1)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	s = "_" + s
	p = "_" + p

	dp[0][0] = true

	for j := 1; j < n+1; j++ {
		if p[j] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {

			if s[i] == p[j] || p[j] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}

			if p[j] == '*' {
				// match zero, make it empty
				dp[i][j] = dp[i][j-2]

				// match one or more times
				if s[i] == p[j-1] || p[j-1] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}

			}

		}
	}

	return dp[m][n]

}
