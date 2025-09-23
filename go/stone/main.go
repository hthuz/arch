package main

import "fmt"

func main() {

	nums := []int{2, 3, -1, 1, 4}
	m := 1
	max_power(m, nums)
}

func max_power(m int, nums []int) int {

	dp := make([]int, len(nums))
	//dp: maximun remaining power jumping at stone i
	for i := 0; i < len(nums); i++ {
		dp[i] = -1e9
	}
	dp[0] = m + nums[0]

	for i := 1; i < len(nums); i++ {
		// for all previous stones
		for j := 0; j < i; j++ {
			// if it can reach position i
			if dp[j] >= i-j {
				// find the max
				dp[i] = max(dp[i], dp[j]-(i-j)+nums[i])
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}
